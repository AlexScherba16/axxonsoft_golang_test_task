package proxy

import (
	"axxonsoft_golang_test_task/config"
	"axxonsoft_golang_test_task/internal/middleware"
	"axxonsoft_golang_test_task/pkg/services/proxy"
	"axxonsoft_golang_test_task/pkg/utils"
	"context"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

type ProxyHandler interface {
	ProxyRequest(w http.ResponseWriter, r *http.Request)
}

type proxyHandler struct {
	requestTimeoutMs time.Duration
	middleware       *middleware.ProxyServerMiddleware
}

func NewProxyHandler(settings *config.ApplicationSettings, mid *middleware.ProxyServerMiddleware) *proxyHandler {
	return &proxyHandler{
		requestTimeoutMs: time.Duration(settings.Timeout.ThirdPartyRequestTimeoutMs) * time.Millisecond,
		middleware:       mid,
	}
}

func (p *proxyHandler) ProxyRequest(w http.ResponseWriter, r *http.Request) {
	var clientRequest proxy.ClientRequest

	// unmarshall raw request body to JSON ClientRequest
	err := utils.ReadJSON(w, r, &clientRequest)
	if err != nil {
		p.sendJsonAndFailIfTransitionError(w, err.Error())
		return
	}

	// check for successful ClientRequest parsing result
	if proxy.CompareRequestStructs(&clientRequest, &proxy.ClientRequest{}) == true {
		p.sendJsonAndFailIfTransitionError(w, "user request is empty")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), p.requestTimeoutMs)
	defer cancel()

	// create raw http request from ClientRequest data structure
	proxyRequest, err := p.generateUserRequest(ctx, &clientRequest)
	if err != nil {
		p.sendJsonAndFailIfTransitionError(w, err.Error())
		return
	}

	// exec request
	client := &http.Client{}
	response, err := client.Do(proxyRequest)
	if err != nil {
		p.sendJsonAndFailIfTransitionError(w, err.Error())
		return
	}
	defer response.Body.Close()

	// unmarshall raw response body to JSON ClientResponse
	clientResponse := p.generateUserResponse(response)
	err = utils.WriteJSON(w, http.StatusOK, clientResponse)
	if err != nil {
		p.sendJsonAndFailIfTransitionError(w, err.Error())
		return
	}

	// update map storage
	err = p.middleware.Update(clientRequest, clientResponse)
	if err != nil {
		p.sendJsonAndFailIfTransitionError(w, err.Error())
		return
	}
	log.Infof("update map with request   : %v", clientRequest)
	log.Infof("update map with reesponse : %v\n", clientResponse)
}

// sendJsonAndFailIfError - generate JSON error message and send via http, shutdown app if transition error occurred
func (p *proxyHandler) sendJsonAndFailIfTransitionError(w http.ResponseWriter, errorString string) {
	log.Error(errorString)
	err := utils.WriteJSON(w, http.StatusBadRequest, errorString)
	if err != nil {
		log.Fatal(err)
	}
}

// generateUserRequest - create raw http.Request from ClientRequest struct
func (p *proxyHandler) generateUserRequest(ctx context.Context, client *proxy.ClientRequest) (*http.Request, error) {
	request, err := http.NewRequestWithContext(ctx, client.Method, client.Url, nil)
	if err != nil {
		return nil, err
	}

	// set custom headers to third party request
	for key, value := range client.Headers {
		request.Header.Set(key, value)
	}
	return request, nil
}

// generateUserResponse - create ClientResponse form raw http.Response struct
func (p *proxyHandler) generateUserResponse(response *http.Response) proxy.ClientResponse {
	clientResponse := proxy.ClientResponse{
		ID:      uuid.NewV4().String(),
		Status:  response.StatusCode,
		Length:  response.ContentLength,
		Headers: make(map[string]string),
	}

	// cast http.Header to map[string]string
	for key, value := range response.Header {
		clientResponse.Headers[key] = strings.Join(value, " ")
	}

	return clientResponse
}
