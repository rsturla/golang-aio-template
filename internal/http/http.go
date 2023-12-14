package http

import (
	"embed"
	"net/http"
)

type Server struct {
	Router        *http.ServeMux
	WebFilesystem embed.FS
}

func NewServer(filesystem embed.FS) *Server {
	return &Server{
		Router:        http.NewServeMux(),
		WebFilesystem: filesystem,
	}
}

func (s *Server) Routes() {
	s.Router.HandleFunc("/", s.handleWeb(s.WebFilesystem))
	s.Router.HandleFunc("/api", s.handleAPI())
}

func (s *Server) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, s.Router)
}
