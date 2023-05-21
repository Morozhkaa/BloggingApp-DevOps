package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Config struct {
	HTTP_port int    `env:"HTTP_PORT" envDefault:"3030"`
	IsProd    bool   `env:"IS_PROD" envDefault:"false"`
	Salt      string `env:"SALT" envDefault:"b9PDPbt4"`
	Secret    string `env:"SECRET" envDefault:"secret"`
	DbUrl     string `env:"DB_URL" envDefault:"postgres://postgres:qwerty@auth-db:5432/auth?sslmode=disable"`
}

var config Config = Config{}

func GetConfig() (*Config, error) {
	if err := env.Parse(&config); err != nil {
		return nil, fmt.Errorf("read logger configuration failed: %w", err)
	}

	return &config, nil
}
