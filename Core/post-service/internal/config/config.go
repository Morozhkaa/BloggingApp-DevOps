package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Config struct {
	HTTP_port int    `env:"HTTP_PORT" envDefault:"9000"`
	IsProd    bool   `env:"IS_PROD" envDefault:"false"`
	AuthURL   string `env:"AUTH_URL" envDefault:"http://localhost:3030/"`
	DbUrl     string `env:"DB_URL" envDefault:"postgres://postgres:qwerty@post-db:5432/post?sslmode=disable"`
}

var config Config = Config{}

func GetConfig() (*Config, error) {
	if err := env.Parse(&config); err != nil {
		return nil, fmt.Errorf("read logger configuration failed: %w", err)
	}

	return &config, nil
}
