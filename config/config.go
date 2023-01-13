package config

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var LocalIp = "127.0.0.1"
var LocalPort = "3384"

var LocalIp2 = "0.0.0.0"
var LocalPort2 = "3456"
var Default_config = Config{
	LogLevel:           "info",
	NumberOfCids:       15,
	HttpServerHostname: "localhost",
	HttpServerPort:     7000,
}

type Config struct {
	LogLevel           string `json:"log-level"`
	NumberOfCids       int    `json:"cid-number"`
	HttpServerHostname string `json:"hostname"`
	HttpServerPort     int    `json:"port"`
}

func NewConfig(Cctx *cli.Context) (*Config, error) {
	c := &Config{}
	switch {
	case Cctx.Command.Name == "run_optimistic_provide":
		if Cctx.IsSet("log-level") {
			c.LogLevel = Cctx.String("log-level")
		} else {
			c.LogLevel = Default_config.LogLevel
		}
		if Cctx.IsSet("cid-number") {
			c.NumberOfCids = Cctx.Int("cid-number")
		} else {
			c.NumberOfCids = Default_config.NumberOfCids
		}
		if Cctx.IsSet("http-port") {
			c.HttpServerPort = Cctx.Int("http-port")
		} else {
			c.HttpServerPort = Default_config.HttpServerPort
		}
		if Cctx.IsSet("http-hostname") {
			c.HttpServerHostname = Cctx.String("http-hostname")
		} else {
			c.HttpServerHostname = Default_config.HttpServerHostname
		}

	case Cctx.Command.Name == "run_cid_hoarder":
		log.Info("run cid hoarder is not implemented yet")
	default:
		return nil, errors.New("unknown command found")
	}
	return c, nil
}
