package config

var local_ip = "127.0.0.1"
var local_port = "8934"

var default_config = Config{
	log_level:      "info",
	number_of_cids: 15,
}

type Config struct {
	log_level      string `json:"log-level"`
	number_of_cids int    `json:"cid-number"`
}
