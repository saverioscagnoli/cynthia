package api

import (
	"bytes"
	"cynthia/ds/dstypes"
	"encoding/json"
	"fmt"
	"net/http"
)

const apiLink = "https://discord.com/api"

func SendMessageContent(token string, channelID dstypes.Snowflake, content string) error {
	url := EndpointCreateMessage(channelID)
	body, err := json.Marshal(map[string]any{"content": content})

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+token)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected status when sending a message: %s", res.Status)
	}

	return nil
}
