package main

import (
	"context"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
	"net/http"
	"os"
	"retrieval-success-rate/config"
	"retrieval-success-rate/pkg"
)

var run_optimistic_provide = &cli.Command{
	Name:   "run_optimistic_provide",
	Usage:  "starts providing cids to the network using optimistic provide",
	Action: provide,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "log-level",
			Usage: "log level [debug,warn,info,error]",
		},
		&cli.IntFlag{
			Name:  "cid-number",
			Usage: "number of cids to provide",
		},
	},
}

var run_cid_hoarder = &cli.Command{
	Name:  "run_cid_hoarder",
	Usage: "starts pinging the providers of the cids in order to gather information about them",
}

func main() {
	log.Info("Starting retrieval success rate")
	retrieval_success_rate := cli.App{
		Name:  "retrieval_success_rate",
		Usage: "publishing data using dennis-tra optimistic-provide and measuring their retrieval success rate using cortze's cid hoarder",
		Commands: []*cli.Command{
			run_optimistic_provide,
			run_cid_hoarder,
		},
	}
	if err := retrieval_success_rate.RunContext(context.Background(), os.Args); err != nil {
		log.Errorf("An error occured while trying to run the app: %v", err)
		os.Exit(1)
	}
}

func provide(Cctx *cli.Context) error {
	log.Info("Starting the provide process")
	new_config_instance, err := config.NewConfig(Cctx)

	log.SetLevel(config.ParseLogLevel(new_config_instance.LogLevel))

	if err != nil {
		return errors.Wrap(err, " error while trying to generate config")
	}

	log.Debugf("Number of cids to provide is: %d", new_config_instance.NumberOfCids)
	log.Debugf("Log level is set to: %s", new_config_instance.LogLevel)
	//what is this?
	go func() {
		profAddr := config.LocalIp + ":" + config.LocalPort
		log.Debugf("Initializing http listen and serve: %s", profAddr)
		err := http.ListenAndServe(profAddr, nil)
		if err != nil {
			log.Errorf("Error initiliazing prometheus at %s with error %s", profAddr, err.Error())
		}
	}()
	//TODO is generating priv key needed?
	host, err := pkg.NewHost(Cctx.Context, config.LocalIp2, config.LocalPort2)

	defer func(host *pkg.Host) {
		err := host.Close()
		if err != nil {
			log.Errorf("error while trying to shut down host")
		}
	}(host)

	if err != nil {
		return errors.Wrap(err, " error while trying to create host")
	}
	err = host.Bootstrap(Cctx.Context)
	if err != nil {
		return errors.Wrap(err, " error while bootstraping the host")
	}

	error_count := 0
	//force the refresh of all buckets
	//<-host.DHT.ForceRefresh()

	for i := 0; i < new_config_instance.NumberOfCids; i++ {
		err := pkg.StartProvidingEstimator(Cctx.Context, host)
		if err != nil {
			error_count = error_count + 1
			log.Errorf("unable to provide cid: %d", i)
		}
		log.Debugf("provided %d cid", i)
	}

	if error_count == new_config_instance.NumberOfCids {
		return errors.New("was unable to provide any cids to the network")
	}

	return nil
}
