package _map

import (
	"axxonsoft_golang_test_task/pkg/services/proxy"
	"errors"
	"github.com/mitchellh/hashstructure"
)

type RequestResponseMap struct {
	responses map[uint64]proxy.ClientResponse
}

func NewRequestResponseMap() *RequestResponseMap {
	return &RequestResponseMap{
		responses: make(map[uint64]proxy.ClientResponse),
	}
}

func (m *RequestResponseMap) Update(req proxy.ClientRequest, resp proxy.ClientResponse) error {
	hash, err := hashstructure.Hash(req, nil)
	if err != nil {
		return err
	}
	m.responses[hash] = resp
	return nil
}

func (m *RequestResponseMap) GetResponse(req proxy.ClientRequest) (error, proxy.ClientResponse) {
	hash, err := hashstructure.Hash(req, nil)
	if err != nil {
		return err, proxy.ClientResponse{}
	}

	response, ok := m.responses[hash]
	if ok {
		return nil, response
	}
	return errors.New("Empty map"), proxy.ClientResponse{}
}
