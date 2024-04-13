package config

import (
	"github.com/hamillka/avitoTech24/internal/db"
	"github.com/hamillka/avitoTech24/internal/logger"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Log  logger.LogConfig  `envconfig:"LOG"`
	DB   db.DatabaseConfig `envconfig:"DB"`
	Port string            `envconfig:"PORT"`
}

func New() (*Config, error) {
	var config Config

	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
