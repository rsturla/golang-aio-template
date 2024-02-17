package cmd

import (
	"embed"
	"fmt"
	"os"

	"github.com/rsturla/golang-aio/internal/config"
	"github.com/rsturla/golang-aio/pkg/log"
)

// EmbedFS must be passed in from the main package since go:embed cannot view
// objects in parent directories
var embedFS embed.FS
var cfg *config.Config

const configEnvPrefix = "GOLANG_AIO_"

func Execute(embed embed.FS) error {
	embedFS = embed

	if err := setupLogger(); err != nil {
		return err
	}
	if err := setupConfig(); err != nil {
		return err
	}

	// Register commands to the application
	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	environment := os.Getenv("ENVIRONMENT")
	log.Debugf("Setting up logger for environment: %s", environment)

	switch environment {
	case "development":
		log.SetFormatter(log.TextFormatter)
		return log.SetLogLevel("debug")
	default:
		return log.SetLogLevel("info")
	}
}

func setupConfig() error {
	configFileEnvName := fmt.Sprintf("%sCONFIG_FILE", configEnvPrefix)
	configFileName := os.Getenv(configFileEnvName)
	log.Debugf("Setting up config with file: %s", configFileName)

	c := config.New()
	if err := c.Load(configFileName, configEnvPrefix); err != nil {
		return err
	}

	cfg = c
	return nil
}
