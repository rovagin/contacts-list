package main

import (
	"github.com/caarlos0/env/v6"
	"time"
)

type Config struct {
	PG   PG
	HTTP HTTP
}

type HTTP struct {
	URI string `env:"HTTP_SERVER_URI" enfDefault:"0.0.0.0:80"`
}

type PG struct {
	URI               string        `env:"PG_URI"`
	ConnectionTimeout time.Duration `env:"PG_CONNECT_TIMEOUT"`
}

func parse() (*Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
