package test

import (
	"axxonsoft_golang_test_task/config"
	"axxonsoft_golang_test_task/internal/middleware"
	"axxonsoft_golang_test_task/internal/services/proxy"
	_map "axxonsoft_golang_test_task/pkg/map"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type handlerTestSuite struct {
	name       string
	body       map[string]interface{}
	statusCode int
}

func TestProxyRequestHandler(t *testing.T) {
	appConfig := config.NewApplicationSettings()
	reqRespMap := _map.NewRequestResponseMap()
	proxyMiddleware := middleware.NewProxyServerMiddleware(reqRespMap)
	proxyHandler := proxy.NewProxyHandler(appConfig, proxyMiddleware)

	suites := generateProxyHandlerTestSuites()
	for _, suite := range suites {
		body, _ := json.Marshal(suite.body)
		request, _ := http.NewRequest("GET", "/proxy/", bytes.NewReader(body))
		response := httptest.NewRecorder()

		handler := http.HandlerFunc(proxyHandler.ProxyRequest)
		handler.ServeHTTP(response, request)
		if response.Code != suite.statusCode {
			t.Fatalf("[%s] expected status: %d result status : %d", suite.name, suite.statusCode, response.Code)
		}
	}
}

func generateProxyHandlerTestSuites() []handlerTestSuite {
	return []handlerTestSuite{
		{name: "failed_raw_request_no_headers", statusCode: http.StatusBadRequest, body: map[string]interface{}{
			"email": "me@here.com",
		}},

		{name: "failed_raw_request_with_headers", statusCode: http.StatusBadRequest, body: map[string]interface{}{
			"email": "me@here.com",
			"headers": map[string]string{
				"X-Amzn-Trace-Id": "*", "X-Forwarded-For": "*.*.*.*",
			},
		}},

		{name: "valid_raw_request_no_headers", statusCode: http.StatusOK, body: map[string]interface{}{
			"method": "GET", "url": "http://httpbin.org/get",
		}},

		{name: "valid_raw_request_with_headers", statusCode: http.StatusOK, body: map[string]interface{}{
			"method": "GET", "url": "http://httpbin.org/get",
			"headers": map[string]string{
				"Authentication":           "Basic bG9naW46cGFzc3dvcmQ=",
				"Sec-WebSocket-Extensions": "permessage-deflate; client_max_window_bits",
			},
		}},

		{name: "valid_raw_google_request_with_headers", statusCode: http.StatusOK, body: map[string]interface{}{
			"method": "GET", "url": "http://google.com",
			"headers": map[string]string{
				"Authentication": "dccGFzc3dvcmQ=",
			},
		}},

		{name: "valid_raw_post_request_without_headers", statusCode: http.StatusOK, body: map[string]interface{}{
			"method": "POST", "url": "http://httpbin.org/post",
		}},

		// add here future requests
	}
}
