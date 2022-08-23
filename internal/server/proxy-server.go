package server

import (
	"axxonsoft_golang_test_task/config"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type proxyServer struct {
	httpServer *http.Server
}

func NewProxyServer(cfg *config.ApplicationSettings) *proxyServer {
	return &proxyServer{
		httpServer: &http.Server{
			Addr:           fmt.Sprintf(":%d", cfg.Server.Port),
			MaxHeaderBytes: cfg.Server.MaxHeaderBytes,
			ReadTimeout:    time.Duration(cfg.Timeout.ServerReadTimeoutMs) * time.Millisecond,
			WriteTimeout:   time.Duration(cfg.Timeout.ServerWriteTimeoutMs) * time.Millisecond,
		},
	}
}

func (s *proxyServer) RunServer(handler http.Handler) error {
	if s.httpServer == nil {
		return errors.New("nil server")
	}
	s.httpServer.Handler = handler
	return s.httpServer.ListenAndServe()
}

func (s *proxyServer) ShutdownServer(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
