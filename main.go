package main

import (
	"embed"
	"fmt"
	"github.com/rsturla/golang-aio/internal/config"
	"github.com/rsturla/golang-aio/internal/http"
	"github.com/rsturla/golang-aio/pkg/log"
	"os"
)

var (
	configEnvPrefix = "GOLANG_AIO_"
)

//go:embed all:web/dist
var embedFS embed.FS

func main() {
	if err := run(os.Args); err != nil {
		log.Fatalf("Run failed with error: %s", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	// Set logging format based on environment.
	if err := setupLogger(); err != nil {
		return err
	}

	// Setup configuration.
	configFileEnvName := fmt.Sprintf("%sCONFIG_FILE", configEnvPrefix)
	cfg, err := setupConfig(os.Getenv(configFileEnvName), configEnvPrefix)
	if err != nil {
		return err
	}

	// Register the API endpoints.
	s := http.NewServer(embedFS, cfg)
	s.Routes()

	// Start the HTTP server.
	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Infof("Starting HTTP server on port %s", addr)
	return s.ListenAndServe(addr)
}

func setupLogger() error {
	switch os.Getenv("ENVIRONMENT") {
	case "development":
		log.SetFormatter(log.TextFormatter)
		return log.SetLogLevel("debug")
	default:
		return log.SetLogLevel("info")
	}
}

func setupConfig(configFilePath, configEnvPrefix string) (*config.Config, error) {
	cfg := config.New()
	if err := cfg.Load(configFilePath, configEnvPrefix); err != nil {
		return nil, err
	}
	return cfg, nil
}
