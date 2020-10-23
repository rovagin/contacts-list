package main

import (
	"contacts-list/internal/pkg/connector"
	"contacts-list/internal/pkg/mongo"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Mongo     mongo.Config
	HTTP      HTTP
	Connector connector.Config
}

type HTTP struct {
	URI string `env:"HTTP_SERVER_URI" enfDefault:"0.0.0.0:8000"`
}

func parse() (*Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
