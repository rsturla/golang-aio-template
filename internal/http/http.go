package http

import (
	"embed"
	"encoding/json"
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

func encode[T any](w http.ResponseWriter, status int, data T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		return fmt.Errorf("failed to encode data: %v", err)
	}

	return nil
}

func decode[T any](r *http.Request) (T, error) {
	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, fmt.Errorf("failed to decode data: %v", err)
	}

	return data, nil
}
