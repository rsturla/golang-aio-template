package http

import (
	"github.com/rsturla/golang-aio/pkg/log"
	"net/http"
)

// HandleAPI is the handler for the dummy API endpoint.
func (s *Server) handleCountAPI() http.HandlerFunc {
	count := 0

	return func(w http.ResponseWriter, r *http.Request) {
		count++
		log.Infof("Count: %d", count)
		if err := encode(w, http.StatusOK, map[string]int{"count": count}); err != nil {
			log.Errorf("Error encoding response: %v", err)
		}
	}
}
