package http

import (
	"embed"
	"fmt"
	"net/http"
	"time"

	"github.com/rsturla/golang-aio/internal/config"
)

type Server struct {
	HttpServer    *http.Server
	Router        *http.ServeMux
	WebFilesystem *embed.FS
	Config        *config.Config
}

func NewServer(filesystem *embed.FS, cfg *config.Config) *Server {
	s := &Server{
		WebFilesystem: filesystem,
		Config:        cfg,
		Router:        http.NewServeMux(),
	}

	s.HttpServer = &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 3,
		IdleTimeout:  time.Second * 60,
		Handler:      s.Router,
	}

	s.setRoutes()
	return s
}
