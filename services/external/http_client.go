package external

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type IHttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type HttpClient struct {
	Client IHttpClient
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		Client: &http.Client{},
	}
}

func (httpClient *HttpClient) Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	if err != nil {
		return nil, err
	}
	request.Header = headers
	return httpClient.Client.Do(request)
}
