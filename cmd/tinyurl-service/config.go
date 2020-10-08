package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	Http struct {
		Port uint16 `default:"8080"`
	}
	Log struct {
		Level string `default:"debug"`
	}
}

func (cfg Config) ListenerAddress() string {
	return fmt.Sprintf("0.0.0.0:%d", cfg.Http.Port)
}

func (cfg Config) CreateLogger() (*zap.Logger, error) {
	return zap.NewDevelopment()
}

func ParseConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("", &config); err != nil {
		return config, err
	}

	return config, nil
}
