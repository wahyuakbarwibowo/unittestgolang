package tokenization

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPost(t *testing.T) {
	mockClient := &MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			body := io.NopCloser(strings.NewReader(`{"message":"success"}`))
			return &http.Response{
				StatusCode: 201,
				Body:       body,
				Header:     make(http.Header),
			}, nil
		},
	}

	client := NewClient(mockClient)
	data := []byte(`{"name": "test"}`)
	resp, err := client.Post("http://example.com", data)

	assert.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, `{"message":"success"}`, string(body))
}
