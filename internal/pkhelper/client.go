package pkhelper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type PokemonClient struct {
	baseURL string
	http    *http.Client
}

var client = newClient("http://localhost:9000")

func newClient(url string) *PokemonClient {
	return &PokemonClient{baseURL: url, http: &http.Client{}}
}

func SetDefaultClient(url string) {
	client = newClient(url)
}

func (c *PokemonClient) request(method string, endpoint string, body any) (*http.Response, error) {
	var req *http.Request
	var err error

	if body != nil {
		b, err := json.Marshal(body)

		if err != nil {
			return nil, fmt.Errorf("Failed to marshal request body: %w", err)
		}

		req, err = http.NewRequest(method, c.baseURL+endpoint, bytes.NewBuffer(b))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(method, c.baseURL+endpoint, nil)
	}

	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %w", err)
	}

	return c.http.Do(req)
}

func (c *PokemonClient) get(endpoint string) (*http.Response, error) {
	return c.request(http.MethodGet, endpoint, nil)
}
