package http

import (
	"github.com/rsturla/golang-aio/pkg/log"
	"net/http"
)

// handleCountAPI is the handler for the dummy API endpoint.
func (s *Server) handleCountAPI() http.HandlerFunc {

	// Define the request and response types for the endpoint
	type response struct {
		Count int `json:"count"`
	}

	// Prepare the handler
	count := 0

	return func(w http.ResponseWriter, r *http.Request) {
		// Increment the count
		count++

		// Prepare and send the response
		response := response{Count: count}
		if err := encode(w, http.StatusOK, response); err != nil {
			log.Errorf("Error encoding response: %v", err)
		}
	}
}
