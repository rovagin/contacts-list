package main

import (
	"github.com/caarlos0/env/v6"
	"time"
)

type Config struct {
	PG PG
}

type PG struct {
	URI string `env:"PG_URI"`
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

