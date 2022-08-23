package test

import (
	"axxonsoft_golang_test_task/config"
	"axxonsoft_golang_test_task/internal/api"
	"axxonsoft_golang_test_task/internal/middleware"
	_map "axxonsoft_golang_test_task/pkg/map"
	"github.com/go-chi/chi/v5"
	"net/http"
	"testing"
)

func TestRoutesExist(t *testing.T) {
	appConfig := config.NewApplicationSettings()
	reqRespMap := _map.NewRequestResponseMap()
	proxyMiddleware := middleware.NewProxyServerMiddleware(reqRespMap)
	appHandler := api.NewApplicationHandler(appConfig, proxyMiddleware)

	testRoutes := appHandler.NewRouters()
	chiRoutes := testRoutes.(chi.Router)

	// TODO: add here future routes
	routes := []string{"/proxy/"}
	for _, route := range routes {
		routeExists(t, chiRoutes, route)
	}
}

func routeExists(t *testing.T, routes chi.Router, route string) {
	found := false

	_ = chi.Walk(routes, func(method string, foundRoute string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if route == foundRoute {
			found = true
		}
		return nil
	})

	if !found {
		t.Errorf("did not find %s in registered routes", route)
	}
}
