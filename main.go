package main

import (
	"embed"
	"fmt"
	"github.com/rsturla/golang-aio/internal/http"
	"github.com/rsturla/golang-aio/pkg/log"
	"github.com/sirupsen/logrus"
	"os"
)

//go:embed all:web/dist
var embedFS embed.FS

func main() {
	// Set logging format based on environment.
	if os.Getenv("ENVIRONMENT") == "development" {
		log.SetLogLevel(logrus.DebugLevel)
		log.SetFormatter(&logrus.TextFormatter{})
	} else {
		log.SetLogLevel(logrus.InfoLevel)
	}

	// Register the API endpoints.
	s := http.NewServer(embedFS)
	s.Routes()

	// Start the HTTP server.
	addr := ":8080"
	serverURL := fmt.Sprintf("http://localhost%s", addr)
	log.Printf("Starting HTTP server at %s ...", serverURL)
	log.Fatal(s.ListenAndServe(addr))
}
