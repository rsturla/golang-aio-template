package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"runtime/pprof"

	"github.com/rsturla/golang-aio/pkg/log"
	"github.com/sirupsen/logrus"
)

//go:embed all:ui/dist
var embedFS embed.FS

func main() {
	if os.Getenv("ENVIRONMENT") == "development" {
		log.SetLogLevel(logrus.DebugLevel)
		log.SetFormatter(&logrus.TextFormatter{})
	} else {
		log.SetLogLevel(logrus.InfoLevel)
	}

	// Extract the embedded filesystem for the frontend.
	distFS, err := fs.Sub(embedFS, "ui/dist")
	if err != nil {
		log.Fatal("Failed to initialize embedded filesystem:", err)
	}

	// Set up HTTP handlers for serving the frontend and handling the API.
	http.Handle("/", http.FileServer(http.FS(distFS)))
	http.HandleFunc("/api", handleAPI)

	// Start the HTTP server.
	addr := ":8080"
	serverURL := fmt.Sprintf("http://localhost%s", addr)
	log.Printf("Starting HTTP server at %s ...", serverURL)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// handleAPI is the handler for the dummy API endpoint.
func handleAPI(w http.ResponseWriter, _ *http.Request) {
	log.Debug("API Endpoint hit")
	if err := writeAllocsProfile(w); err != nil {
		log.Printf("Error: Failed to write allocs profile: %v", err)
	}
}

// writeAllocsProfile writes the allocs profile to the HTTP response.
func writeAllocsProfile(w http.ResponseWriter) error {
	// Retrieve the allocs profile.
	profile := pprof.Lookup("allocs")

	// Write the allocs profile (human-readable, via debug: 1) to the HTTP response.
	return profile.WriteTo(w, 1)
}
