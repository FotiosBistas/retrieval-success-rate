package config

import (
	"errors"
	"github.com/urfave/cli/v2"
)

var Local_ip = "127.0.0.1"
var Local_port = "8934"

var Default_config = Config{
	Log_level:      "info",
	Number_of_cids: 15,
}

type Config struct {
	Log_level      string `json:"log-level"`
	Number_of_cids int    `json:"cid-number"`
}

func NewConfig(Cctx *cli.Context) (*Config, error) {
	c := &Config{}
	switch {
	case Cctx.Command.Name == "optimistic_provide":
		if Cctx.IsSet("log-level") {
			c.Log_level = Cctx.String("log-level")
		}
		if Cctx.IsSet("cid-number") {
			c.Number_of_cids = Cctx.Int("cid-number")
		}
	default:
		return nil, errors.New("Unknown command found")
	}
	return c, nil
}
