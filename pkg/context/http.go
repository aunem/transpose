package context

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// HTTPRequest is passed to a plugins request method
type HTTPRequest struct {
	ID      string
	Meta    map[string]string
	Request *http.Request
	RW      http.ResponseWriter
}

// HTTPResponse is passed to a plugins response method
type HTTPResponse struct {
	ID       string
	Meta     map[string]string
	Request  *http.Request
	Response *http.Response
	RW       http.ResponseWriter
}

// NewHTTPRequest returns a new request context
func NewHTTPRequest(req *http.Request) *HTTPRequest {
	u := uuid.NewV4()
	return &HTTPRequest{
		ID:      u.String(),
		Request: req,
	}
}

// GetID returns the ID
func (req *HTTPRequest) GetID() string {
	return req.ID
}

// GetMeta returns the request metadata
func (req *HTTPRequest) GetMeta() map[string]string {
	return req.Meta
}

// GetRequest returns the request
func (req *HTTPRequest) GetRequest() interface{} {
	return req.Request
}

// NewHTTPResponse returns a new response context
func NewHTTPResponse(req *HTTPRequest, resp *http.Response) *HTTPResponse {
	return &HTTPResponse{
		ID:       req.ID,
		Request:  req.Request,
		Response: resp,
	}
}

// GetID returns the ID
func (resp *HTTPResponse) GetID() string {
	return resp.ID
}

// GetMeta returns the response metadata
func (resp *HTTPResponse) GetMeta() map[string]string {
	return resp.Meta
}

// GetRequest returns the request
func (resp *HTTPResponse) GetRequest() interface{} {
	return resp.Request
}

// GetResponse returns the response
func (resp *HTTPResponse) GetResponse() interface{} {
	return resp.Response
}

// RequestToHTTP takes a request interface and attempts to cast it as an HTTPRequest
func RequestToHTTP(request Request) (*HTTPRequest, error) {
	req, ok := request.(*HTTPRequest)
	if !ok {
		return nil, fmt.Errorf("could not cast requst as HTTPRequest type")
	}
	return req, nil
}

// ResponseToHTTP takes a request interface and attempts to cast it as an HTTPResponse
func ResponseToHTTP(response Response) (*HTTPResponse, error) {
	resp, ok := response.(*HTTPResponse)
	if !ok {
		return nil, fmt.Errorf("could not cast response as HTTPResponse type")
	}
	return resp, nil
}
