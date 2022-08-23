package server

import (
	"context"
	"net/http"
)

type ServerInterface interface {
	RunServer(handler http.Handler) error
	ShutdownServer(ctx context.Context) error
}
