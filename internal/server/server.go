package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(
	address string,
	handler http.Handler,
	readTimeout time.Duration,
	writeTimeout time.Duration,
) error {
	s.httpServer = &http.Server{
		Addr:           address,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 Mb
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
