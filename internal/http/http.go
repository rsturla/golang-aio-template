package http

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/rsturla/golang-aio/internal/config"
	"github.com/rsturla/golang-aio/pkg/log"
)

type Server struct {
	Router        *http.ServeMux
	WebFilesystem embed.FS
	Config        *config.Config
}

func Serve(filesystem embed.FS, cfg *config.Config) error {
	s := newServer(filesystem, cfg)
	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Infof("Listening on %s", addr)
	return s.listenAndServe(addr)
}

func newServer(filesystem embed.FS, cfg *config.Config) *Server {
	s := &Server{
		Router:        http.NewServeMux(),
		WebFilesystem: filesystem,
		Config:        cfg,
	}

	s.setRoutes()
	return s
}

func (s *Server) listenAndServe(addr string) error {
	return http.ListenAndServe(addr, s.Router)
}
