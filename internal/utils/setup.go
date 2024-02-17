package utils

import (
	"os"

	"github.com/rsturla/golang-aio/internal/config"
	"github.com/rsturla/golang-aio/pkg/log"
)

func SetupConfig(configFilePath, configEnvPrefix string) (*config.Config, error) {
	cfg := config.New()
	if err := cfg.Load(configFilePath, configEnvPrefix); err != nil {
		return nil, err
	}
	return cfg, nil
}

func SetupLogger() error {
	environment := os.Getenv("ENVIRONMENT")
	log.Infof("Setting up logger for environment: %s", environment)

	switch environment {
	case "development":
		log.SetFormatter(log.TextFormatter)
		return log.SetLogLevel("debug")
	default:
		return log.SetLogLevel("info")
	}
}
