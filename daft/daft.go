package daft

import (
	"bytes"
	"net/http"
)

const (
	DEFAULT_GATEWAY_URL = "https://search-gateway.dsch.ie/v1"
	DAFT_URL            = "https://www.daft.ie"
)

var DEFAULT_HEADERS = http.Header{
	"content-type": []string{"application/json"},
	"brand":        []string{"daft"},
	"platform":     []string{"web"},
}

type Client struct {
	gatewayUrl string
	client     *http.Client
}

func New() Client {
	client := Client{}
	client.Init()
	return client
}

func (cl *Client) Init() {
	cl.client = &http.Client{}
	cl.gatewayUrl = DEFAULT_GATEWAY_URL
}

func (cl *Client) newRequest(method string, path string, body []byte, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, DEFAULT_GATEWAY_URL+path, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	req.Header = DEFAULT_HEADERS
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := cl.client.Do(req)
	if err != nil {
		panic(err)
	}
	return resp, err
}
