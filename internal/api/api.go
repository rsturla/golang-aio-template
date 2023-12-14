// internal/api/api.go
package api

import (
	"net/http"
	"runtime/pprof"

	"github.com/rsturla/golang-aio/pkg/log"
)

// HandleAPI is the handler for the dummy API endpoint.
func HandleAPI(w http.ResponseWriter, _ *http.Request) {
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
