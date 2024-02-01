package main

import (
	"embed"
	"github.com/rsturla/golang-aio/internal/http"
	"github.com/rsturla/golang-aio/pkg/log"
	"os"
)

//go:embed all:web/dist
var embedFS embed.FS

func main() {
	if err := run(os.Args); err != nil {
		log.Errorf("Run failed with error: %s\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	// Set logging format based on environment.
	if err := setupLogger(); err != nil {
		return err
	}

	// Register the API endpoints.
	s := http.NewServer(embedFS)
	s.Routes()

	// Start the HTTP server.
	addr := ":8080"
	log.Printf("Starting HTTP server on port %s ...\n", addr)
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
