package http

import (
	"github.com/rsturla/golang-aio/pkg/log"
	"net/http"
)

// HandleAPI is the handler for the dummy API endpoint.
func (s *Server) handleCountAPI() http.HandlerFunc {
	count := 0

	type response struct {
		Count int `json:"count"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		count++
		log.Infof("Count: %d", count)

		response := response{Count: count}
		if err := encode(w, http.StatusOK, response); err != nil {
			log.Errorf("Error encoding response: %v", err)
		}
	}
}
