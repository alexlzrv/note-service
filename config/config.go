package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		PG `yaml:"postgres"`
	}

	PG struct {
		PoolMax int    `env-required:"true" yaml:"pool_max"`
		URL     string `env-required:"true" yaml:"url"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
