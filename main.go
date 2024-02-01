package main

import (
	"embed"
	"fmt"
	"github.com/rsturla/golang-aio/internal/config"
	"github.com/rsturla/golang-aio/internal/http"
	"github.com/rsturla/golang-aio/pkg/log"
	"os"
)

//go:embed all:web/dist
var embedFS embed.FS

func main() {
	if err := run(os.Args); err != nil {
		log.Errorf("Run failed with error: %s", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	// Set logging format based on environment.
	if err := setupLogger(); err != nil {
		return err
	}

	// Setup configuration.
	cfg, err := setupConfig(os.Getenv("CONFIG_FILE"))
	if err != nil {
		return err
	}

	// Register the API endpoints.
	s := http.NewServer(embedFS, cfg)
	s.Routes()

	// Start the HTTP server.
	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Starting HTTP server on port %s", addr)
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

func setupConfig(filePath string) (*config.Config, error) {
	cfg := config.New()

	if filePath == "" {
		log.Info("No configuration file specified, using defaults")
		return cfg, nil
	}

	if err := cfg.Load(filePath); err != nil {
		return nil, err
	}

	return cfg, nil
}
