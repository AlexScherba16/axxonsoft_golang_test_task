package test

import (
	"axxonsoft_golang_test_task/pkg/services/proxy"
	"testing"
)

const (
	testMethod = "TEST_METHOD"
	testUrl    = "TEST_URL"
	testId     = "TEST_ID"
	testStatus = 123
	testLen    = 321
)

type requestsTestSuite struct {
	name   string
	a      proxy.ClientRequest
	b      proxy.ClientRequest
	result bool
}

type responseTestSuite struct {
	name   string
	a      proxy.ClientResponse
	b      proxy.ClientResponse
	result bool
}

func TestCompareRequestStructs(t *testing.T) {
	suites := []requestsTestSuite{
		{name: "empty_requests", a: proxy.ClientRequest{}, b: proxy.ClientRequest{}, result: true},
		{name: "correct_requests",
			a:      proxy.ClientRequest{Method: testMethod, Url: testUrl, Headers: map[string]string{"key_0": "value_0", "key_1": "value_1"}},
			b:      proxy.ClientRequest{Method: testMethod, Url: testUrl, Headers: map[string]string{"key_0": "value_0", "key_1": "value_1"}},
			result: true},
		{name: "incorrect_request_method",
			a:      proxy.ClientRequest{Url: testUrl, Headers: map[string]string{"key_0": "value_0", "key_1": "value_1"}},
			b:      proxy.ClientRequest{Method: testMethod, Url: testUrl, Headers: map[string]string{"key_0": "value_0", "key_1": "value_1"}},
			result: false},
		{name: "incorrect_request_url",
			a:      proxy.ClientRequest{Method: testMethod, Headers: map[string]string{"key_0": "value_0", "key_1": "value_1"}},
			b:      proxy.ClientRequest{Method: testMethod, Url: testUrl, Headers: map[string]string{"key_0": "value_0", "key_1": "value_1"}},
			result: false},
		{name: "incorrect_request_header_len",
			a:      proxy.ClientRequest{Method: testMethod, Url: testUrl, Headers: map[string]string{"key_0": "value_0"}},
			b:      proxy.ClientRequest{Method: testMethod, Url: testUrl, Headers: map[string]string{"key_0": "value_0", "key_1": "value_1"}},
			result: false},
		{name: "incorrect_request_header_key",
			a:      proxy.ClientRequest{Method: testMethod, Url: testUrl, Headers: map[string]string{"key_12": "value_0", "key_1": "value_1"}},
			b:      proxy.ClientRequest{Method: testMethod, Url: testUrl, Headers: map[string]string{"key_0": "value_0", "key_1": "value_1"}},
			result: false},
		{name: "incorrect_request_header_value",
			a:      proxy.ClientRequest{Method: testMethod, Url: testUrl, Headers: map[string]string{"key_0": "value_123", "key_1": "value_1"}},
			b:      proxy.ClientRequest{Method: testMethod, Url: testUrl, Headers: map[string]string{"key_0": "value_0", "key_1": "value_1"}},
			result: false},
		{name: "incorrect_request_header_key_value",
			a:      proxy.ClientRequest{Method: testMethod, Url: testUrl, Headers: map[string]string{"key_10": "value_123", "key_1": "value_1"}},
			b:      proxy.ClientRequest{Method: testMethod, Url: testUrl, Headers: map[string]string{"key_0": "value_0", "key_1": "value_1"}},
			result: false},
	}
	for _, suite := range suites {
		result := proxy.CompareRequestStructs(&suite.a, &suite.b)
		if result != suite.result {
			t.Fatalf("[%s]: failed, expected : %t, current result : %t", suite.name, suite.result, result)
		}
	}
}

func TestCompareResponseStructs(t *testing.T) {
	suites := []responseTestSuite{
		{name: "empty_responses", a: proxy.ClientResponse{}, b: proxy.ClientResponse{}, result: true},
		{name: "correct_responses",
			a:      proxy.ClientResponse{ID: testId, Status: testStatus, Length: testLen, Headers: map[string]string{"key_0": "value_0", "key_1": "value_1"}},
			b:      proxy.ClientResponse{ID: testId, Status: testStatus, Length: testLen, Headers: map[string]string{"key_0": "value_0", "key_1": "value_1"}},
			result: true},
		{name: "incorrect_response_id",
			a:      proxy.ClientResponse{Status: testStatus, Length: testLen, Headers: map[string]string{"key_0": "value_0", "key_1": "value_1"}},
			b:      proxy.ClientResponse{ID: testId, Status: testStatus, Length: testLen, Headers: map[string]string{"key_0": "value_0", "key_1": "value_1"}},
			result: false},
		{name: "incorrect_response_status",
			a:      proxy.ClientResponse{ID: testId, Status: testStatus + 1, Length: testLen, Headers: map[string]string{"key_0": "value_0", "key_1": "value_1"}},
			b:      proxy.ClientResponse{ID: testId, Status: testStatus, Length: testLen, Headers: map[string]string{"key_0": "value_0", "key_1": "value_1"}},
			result: false},
		{name: "incorrect_response_len",
			a:      proxy.ClientResponse{ID: testId, Status: testStatus, Length: testLen + 1, Headers: map[string]string{"key_0": "value_0", "key_1": "value_1"}},
			b:      proxy.ClientResponse{ID: testId, Status: testStatus, Length: testLen, Headers: map[string]string{"key_0": "value_0", "key_1": "value_1"}},
			result: false},
		{name: "incorrect_response_header_len",
			a:      proxy.ClientResponse{ID: testId, Status: testStatus, Length: testLen, Headers: map[string]string{"key_0": "value_0"}},
			b:      proxy.ClientResponse{ID: testId, Status: testStatus, Length: testLen, Headers: map[string]string{"key_0": "value_0", "key_1": "value_1"}},
			result: false},
		{name: "incorrect_response_header_key",
			a:      proxy.ClientResponse{ID: testId, Status: testStatus, Length: testLen, Headers: map[string]string{"key_12": "value_0", "key_1": "value_1"}},
			b:      proxy.ClientResponse{ID: testId, Status: testStatus, Length: testLen, Headers: map[string]string{"key_0": "value_0", "key_1": "value_1"}},
			result: false},
		{name: "incorrect_response_header_value",
			a:      proxy.ClientResponse{ID: testId, Status: testStatus, Length: testLen, Headers: map[string]string{"key_0": "value_123", "key_1": "value_1"}},
			b:      proxy.ClientResponse{ID: testId, Status: testStatus, Length: testLen, Headers: map[string]string{"key_0": "value_0", "key_1": "value_1"}},
			result: false},
		{name: "incorrect_response_header_key_value",
			a:      proxy.ClientResponse{ID: testId, Status: testStatus, Length: testLen, Headers: map[string]string{"key_10": "value_123", "key_1": "value_1"}},
			b:      proxy.ClientResponse{ID: testId, Status: testStatus, Length: testLen, Headers: map[string]string{"key_0": "value_0", "key_1": "value_1"}},
			result: false},
	}
	for _, suite := range suites {
		result := proxy.CompareResponseStructs(&suite.a, &suite.b)
		if result != suite.result {
			t.Fatalf("[%s]: failed, expected : %t, current result : %t", suite.name, suite.result, result)
		}
	}
}
