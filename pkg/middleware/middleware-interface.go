package middleware

import (
	"axxonsoft_golang_test_task/pkg/services/proxy"
)

// MiddlewareInterface - generic wrapper over an entity that has an interface Update, Get
type MiddlewareInterface interface {
	Update(req proxy.ClientRequest, resp proxy.ClientResponse) error
	Get(req proxy.ClientRequest) (error, proxy.ClientResponse)
}
