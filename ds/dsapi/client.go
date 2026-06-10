package dsapi

import (
	"bytes"
	"cynthia/ds/dstypes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

const ApiURL = "https://discord.com/api"

type Client struct {
	token string
	appID dstypes.Snowflake
	http  *http.Client
}

func NewApiClient(token string, appID dstypes.Snowflake) *Client {
	return &Client{
		token: token,
		appID: appID,
		http:  http.DefaultClient,
	}
}

func (c *Client) Request(method, url string, body any) ([]byte, error) {
	slog.Debug("Performing request", "method", method, "url", url, "body", body)
	var buf io.Reader

	if body != nil {
		b, err := json.Marshal(body)

		if err != nil {
			return nil, err
		}

		buf = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(method, url, buf)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+c.token)

	res, err := c.http.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("discord API error %d: %s", res.StatusCode, string(data))
	}

	return data, nil
}
