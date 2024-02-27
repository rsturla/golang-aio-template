package http

func (s *Server) setRoutes() {
	s.Router.HandleFunc("/", s.handleWeb(s.WebFilesystem))
	s.Router.HandleFunc("/api/count", s.handleCountAPI())
}
