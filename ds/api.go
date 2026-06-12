package ds

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const ApiURL = "https://discord.com/api"
const CdnURL = "https://cdn.discordapp.com"

type routes struct{}

var Routes = routes{}

type ApiClient struct {
	botToken string
	appID    Snowflake
	http     *http.Client
}

func newApiClient(botToken string, appID Snowflake) *ApiClient {
	return &ApiClient{
		botToken: botToken,
		appID:    appID,
		http:     &http.Client{},
	}
}

func (c *ApiClient) request(method string, endpoint string, body any) (*http.Response, error) {
	var req *http.Request
	var err error

	if body != nil {
		b, err := json.Marshal(body)

		if err != nil {
			return nil, fmt.Errorf("Failed to marshal request body: %w", err)
		}

		req, err = http.NewRequest(method, ApiURL+endpoint, bytes.NewBuffer(b))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(method, ApiURL+endpoint, nil)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bot "+c.botToken)

	return c.http.Do(req)
}

func (c *ApiClient) Get(endpoint string) (*http.Response, error) {
	return c.request(http.MethodGet, endpoint, nil)
}

func (c *ApiClient) Post(endpoint string, body any) (*http.Response, error) {
	return c.request(http.MethodPost, endpoint, body)
}

func (c *ApiClient) Patch(endpoint string, body any) (*http.Response, error) {
	return c.request(http.MethodPatch, endpoint, body)
}

func (c *ApiClient) Put(endpoint string, body any) (*http.Response, error) {
	return c.request(http.MethodPut, endpoint, body)
}

func (c *ApiClient) Delete(endpoint string) (*http.Response, error) {
	return c.request(http.MethodDelete, endpoint, nil)
}
