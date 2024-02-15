package http

import (
	"github.com/rsturla/golang-aio/pkg/log"
	"net/http"
	"runtime/pprof"
)

// HandleAPI is the handler for the dummy API endpoint.
func (s *Server) handleAPI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := writeAllocsProfile(w); err != nil {
			log.Errorf("Failed to write allocs profile: %v\n", err)
		}
	}
}

// writeAllocsProfile writes the allocs profile to the HTTP response.
func writeAllocsProfile(w http.ResponseWriter) error {
	// Retrieve the allocs profile.
	profile := pprof.Lookup("allocs")

	// Write the allocs profile (human-readable, via debug: 1) to the HTTP response.
	return profile.WriteTo(w, 1)
}
