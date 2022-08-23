package api

import (
	"axxonsoft_golang_test_task/config"
	"axxonsoft_golang_test_task/internal/middleware"
	"axxonsoft_golang_test_task/internal/services/proxy"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ApplicationHandler struct {
	proxyHandler proxy.ProxyHandler
}

func NewApplicationHandler(settings *config.ApplicationSettings, mid *middleware.ProxyServerMiddleware) *ApplicationHandler {
	return &ApplicationHandler{
		proxyHandler: proxy.NewProxyHandler(settings, mid),
	}
}

func (h *ApplicationHandler) NewRouters() http.Handler {
	mux := chi.NewRouter()

	mux.Route("/proxy", func(r chi.Router) {
		r.Get("/", h.proxyHandler.ProxyRequest)
	})

	return mux
}
