package retrieval_success_rate

import (
	"context"
	"github.com/FotiosBistas/retrieval-success-rate/config"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
	"net/http"
	"os"
)

var run_optimistic_provide = &cli.Command{
	Name:   "optimistic_provide",
	Usage:  "starts providing cids to the network using optimistic provide",
	Action: provide,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "log-level",
			Usage: "log level [debug,warn,info,error]",
			Value: config.Default_config.Log_level,
		},
		&cli.IntFlag{
			Name:  "cid-number",
			Usage: "number of cids to provide",
			Value: config.Default_config.Number_of_cids,
		},
	},
}

func main() {
	retrieval_success_rate := cli.App{
		Name:  "retrieval_succes_rate",
		Usage: "publishing data using dennis-tra optimistic-provide and measuring their retrieval success rate",
		Commands: []*cli.Command{
			run_optimistic_provide,
		},
	}

	if err := retrieval_success_rate.RunContext(context.Background(), os.Args); err != nil {
		log.Errorf("An error occured while trying to run the app: %v", err)
		os.Exit(1)
	}
}

func provide(Cctx *cli.Context) error {

	new_config, err := config.NewConfig(Cctx)

	if err != nil {
		return errors.Wrap(err, " error while trying to generate config")
	}
	//what is this?
	go func() {
		profAddr := config.Local_ip + ":" + config.Local_port
		log.Debugf("Initializing http listen and serve: %s", profAddr)
		err := http.ListenAndServe(profAddr, nil)
		if err != nil {
			log.Errorf("Error initiliazing prometheus at %s with error %s", profAddr, err.Error())
		}
	}()

	return nil
}
