package proxy

// ClientRequest - a structure for storing data received from the client to create a proxy request
type ClientRequest struct {
	Method  string            `json:"method" binding:"required"`
	Url     string            `json:"url" binding:"required"`
	Headers map[string]string `json:"headers"`
}

// ClientResponse - a structure for storing data received from third party service
type ClientResponse struct {
	ID      string            `json:"id" binding:"required"`
	Status  int               `json:"status" binding:"required"`
	Length  int64             `json:"length" binding:"required"`
	Headers map[string]string `json:"headers,omitempty"`
}
