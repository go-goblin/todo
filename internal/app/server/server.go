package server

import (
	"context"
	"net/http"
	"time"
	"url-stortener/internal/dependencies"
	"url-stortener/internal/handler"
)

type HTTPServer struct {
	server *http.Server
}

func New(deps *dependencies.Dependencies) *HTTPServer {
	h := handler.New(deps)
	r := h.GetRouter()
	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
		Handler:           r,
	}
	return &HTTPServer{
		server: server,
	}
}

func (s *HTTPServer) Run() error {
	err := s.server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
