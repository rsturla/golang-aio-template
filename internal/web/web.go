package web

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/rsturla/golang-aio/pkg/log"
)

func HandleWeb(filesystem embed.FS) http.Handler {
	// Extract the embedded filesystem for the frontend.
	distFS, err := fs.Sub(filesystem, "ui/dist")
	if err != nil {
		log.Fatal("Failed to initialize embedded filesystem:", err)
	}

	// Set up HTTP handlers for serving the frontend and handling the API.
	return http.FileServer(http.FS(distFS))
}
