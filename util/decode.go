package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Decode[T any](res *http.Response) (*T, error) {
	defer res.Body.Close()
	var v T

	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return nil, fmt.Errorf("Failed to decode response: %w", err)
	}

	return &v, nil
}
