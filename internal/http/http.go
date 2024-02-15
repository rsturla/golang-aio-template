package http

import (
	"embed"
	"github.com/rsturla/golang-aio/internal/config"
	"net/http"
)

type Server struct {
	Router        *http.ServeMux
	WebFilesystem embed.FS
	Config        *config.Config
}

func NewServer(filesystem embed.FS, cfg *config.Config) *Server {
	s := &Server{
		Router:        http.NewServeMux(),
		WebFilesystem: filesystem,
		Config:        cfg,
	}

	s.setRoutes()
	return s
}

func (s *Server) setRoutes() {
	s.Router.HandleFunc("/", s.handleWeb(s.WebFilesystem))
	s.Router.HandleFunc("/api", s.handleAPI())
}

func (s *Server) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, s.Router)
}
