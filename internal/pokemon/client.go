package pokemon

import (
	"bytes"
	"cynthia/cmd/pkapi"
	"encoding/json"
	"fmt"
	"net/http"
)

type PokemonSprite = pkapi.PokemonSprite
type PokemonSprites = pkapi.PokemonSprites
type Pokemon = pkapi.Pokemon

type PokemonClient struct {
	baseURL string
	http    *http.Client
}

func NewClient(url string) *PokemonClient {
	return &PokemonClient{baseURL: url}
}

var DefaultClient = NewClient("localhost:9000")

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
