package config

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var LocalIp = "127.0.0.1"
var LocalPort = "8934"

var Default_config = Config{
	LogLevel:     "info",
	NumberOfCids: 15,
}

type Config struct {
	LogLevel     string `json:"log-level"`
	NumberOfCids int    `json:"cid-number"`
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
	case Cctx.Command.Name == "run_cid_hoarder":
		log.Info("run cid hoarder is not implemented yet")
	default:
		return nil, errors.New("unknown command found")
	}
	return c, nil
}
