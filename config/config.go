package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DB Mongo
}

type Mongo struct {
	URI      string
	Username string
	Password string
	Database string
}

func New() (*Config, error) {
	cfg := new(Config)
	err := envconfig.Process("db", &cfg.DB)

	return cfg, err
}
