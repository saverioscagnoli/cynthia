package ds

import (
	"camilla/service/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func (r *routes) CreateMessage(channelID Snowflake) (string, string) {
	return http.MethodPost, fmt.Sprintf("/channels/%s/messages", channelID)
}

func (r *routes) EditMessage(channelID Snowflake, messageID Snowflake) (string, string) {
	return http.MethodPatch, fmt.Sprintf("/channels/%s/messages/%s", channelID, messageID)
}

type MessageBody struct {
	Content     string             `json:"content"`
	Embeds      []*Embed           `json:"embeds"`
	TTS         *bool              `json:"tts"`
	Attachments []*Attachment      `json:"attachments"`
	Components  []MessageComponent `json:"components"`
	Files       []*MessageFile     `json:"files"`
}

func (b *MessageBody) MarshalJSON() ([]byte, error) {
	type Alias MessageBody
	components := make([]json.RawMessage, len(b.Components))

	for i, c := range b.Components {
		data, err := json.Marshal(c)
		if err != nil {
			return nil, err
		}
		components[i] = data
	}

	return json.Marshal(&struct {
		*Alias
		Components []json.RawMessage `json:"components"`
	}{
		Alias:      (*Alias)(b),
		Components: components,
	})
}

func (c *ApiClient) SendMessage(channelID Snowflake, msg *MessageBody) (*Message, error) {
	method, endpoint := Routes.CreateMessage(channelID)

	if len(msg.Files) > 0 {
		return c.sendMultipart(method, endpoint, msg)
	}

	res, err := c.request(method, endpoint, msg)

	if err != nil {
		return nil, err
	}

	return util.Decode[Message](res)
}

func (c *ApiClient) sendMultipart(
	method, endpoint string,
	msg *MessageBody,
) (*Message, error) {
	buf, contentType, err := createMultipartBody(msg, msg.Files)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, ApiURL+endpoint, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", "Bot "+c.botToken)

	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	return util.Decode[Message](res)
}

func (c *ApiClient) SendMessageText(channelID Snowflake, content string) (*Message, error) {
	return c.SendMessage(channelID, &MessageBody{Content: content})
}

func (c *ApiClient) SendMessageEmbed(channelID Snowflake, embed *Embed) (*Message, error) {
	return c.SendMessage(channelID, &MessageBody{Embeds: []*Embed{embed}})
}

func (c *ApiClient) EditMessage(channelID Snowflake, messageID Snowflake, body *MessageBody) (*Message, error) {
	method, endpoint := Routes.EditMessage(channelID, messageID)

	if len(body.Files) > 0 {
		return c.sendMultipart(method, endpoint, body)
	}

	res, err := c.request(method, endpoint, body)

	if err != nil {
		return nil, err
	}

	return util.Decode[Message](res)
}

func (c *ApiClient) EditMessageText(channelID Snowflake, messageID Snowflake, text string) (*Message, error) {
	return c.EditMessage(channelID, messageID, &MessageBody{
		Content: text,
	})
}
