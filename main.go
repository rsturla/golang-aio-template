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
	if os.Getenv("ENVIRONMENT") == "development" {
		log.SetLogLevel(log.DebugLevel)
		log.SetFormatter(log.TextFormatter)
	} else {
		log.SetLogLevel(log.InfoLevel)
	}

	// Register the API endpoints.
	s := http.NewServer(embedFS)
	s.Routes()

	// Start the HTTP server.
	addr := ":8080"
	log.Printf("Starting HTTP server on port %s ...\n", addr)
	return s.ListenAndServe(addr)
}
