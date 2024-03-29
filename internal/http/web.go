package http

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/rsturla/golang-aio/pkg/log"
)

func (s *Server) handleWeb(filesystem *embed.FS) http.HandlerFunc {
	// Extract the embedded filesystem for the frontend.
	distFS, err := fs.Sub(filesystem, "web/dist")
	if err != nil {
		log.Fatalf("Failed to initialize embedded filesystem: %v\n", err)
	}

	// Create a file server for the frontend.
	fileServer := http.FileServer(http.FS(distFS))

	return func(w http.ResponseWriter, r *http.Request) {
		fileServer.ServeHTTP(w, r)
	}
}
