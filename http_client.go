package tokenization

import (
	"bytes"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	HTTPClient HTTPClient
}

func NewClient(httpClient HTTPClient) *Client {
	return &Client{
		HTTPClient: httpClient,
	}
}

func (c *Client) Fetch(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.HTTPClient.Do(req)
}

func (c *Client) Post(url string, data []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return c.HTTPClient.Do(req)
}
