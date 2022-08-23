package middleware

import (
	_map "axxonsoft_golang_test_task/pkg/map"
	"axxonsoft_golang_test_task/pkg/services/proxy"
)

type ProxyServerMiddleware struct {
	encapsulatedMap *_map.RequestResponseMap
}

func NewProxyServerMiddleware(controlMap *_map.RequestResponseMap) *ProxyServerMiddleware {
	return &ProxyServerMiddleware{
		encapsulatedMap: controlMap,
	}
}

func (p *ProxyServerMiddleware) Update(req proxy.ClientRequest, resp proxy.ClientResponse) error {
	return p.encapsulatedMap.Update(req, resp)
}

func (p *ProxyServerMiddleware) Get(req proxy.ClientRequest) (error, proxy.ClientResponse) {
	return p.encapsulatedMap.GetResponse(req)
}
