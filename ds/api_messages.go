package ds

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
)

func (r *routes) CreateMessage(channelID Snowflake) (string, string) {
	return http.MethodPost, fmt.Sprintf("/channels/%s/messages", channelID)
}

type CreateMessageBody struct {
	Content     string             `json:"content"`
	Embeds      []*Embed           `json:"embeds"`
	TTS         *bool              `json:"tts"`
	Attachments []*Attachment      `json:"attachments"`
	Components  []MessageComponent `json:"components"`
	Files       []*MessageFile     `json:"files"`
}

func (c *ApiClient) SendMessage(channelID Snowflake, msg *CreateMessageBody) error {
	method, endpoint := Routes.CreateMessage(channelID)

	if len(msg.Files) > 0 {
		return c.sendMultipart(method, endpoint, msg)
	}

	return c.sendJSON(method, endpoint, msg)
}

func (c *ApiClient) sendJSON(method, endpoint string, msg *CreateMessageBody) error {
	_, err := c.request(method, endpoint, msg)
	return err
}

func (c *ApiClient) sendMultipart(method, endpoint string, msg *CreateMessageBody) error {
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)

	// write payload_json
	j, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	if err := writer.WriteField("payload_json", string(j)); err != nil {
		return err
	}

	// write each file
	for i, f := range msg.Files {
		ct := f.ContentType
		if ct == "" {
			ct = "application/octet-stream"
		}

		part, err := writer.CreatePart(textproto.MIMEHeader{
			"Content-Disposition": {fmt.Sprintf(`form-data; name="files[%d]"; filename="%s"`, i, f.Name)},
			"Content-Type":        {ct},
		})

		if err != nil {
			return err
		}

		if _, err := io.Copy(part, f.Reader); err != nil {
			return err
		}
	}

	writer.Close()

	req, err := http.NewRequest(method, ApiURL+endpoint, buf)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bot "+c.botToken)

	res, err := c.http.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("discord error %d: %s", res.StatusCode, body)
	}

	return nil
}

func (c *ApiClient) SendMessageText(channelID Snowflake, content string) error {
	return c.SendMessage(channelID, &CreateMessageBody{Content: content})
}

func (c *ApiClient) SendMessageEmbed(channelID Snowflake, embed *Embed) error {
	return c.SendMessage(channelID, &CreateMessageBody{Embeds: []*Embed{embed}})
}
