package dsapi

import (
	"cynthia/ds/dstypes"
	"fmt"
)

func EndpointCreateMessage(channelID dstypes.Snowflake) string {
	return fmt.Sprintf("%s/channels/%s/messages", ApiURL, channelID)
}

type CreateMessageBody struct {
	Content     string                `json:"content"`
	Embeds      *[]dstypes.Embed      `json:"embeds"`
	TTS         bool                  `json:"tts"`
	Attachments *[]dstypes.Attachment `json:"attachments"`
}

func (c *Client) SendMessage(channelID dstypes.Snowflake, body CreateMessageBody) error {
	_, err := c.Request("POST", EndpointCreateMessage(channelID), body)
	return err

}

func (c *Client) SendMessageText(channelID dstypes.Snowflake, text string) error {
	body := CreateMessageBody{
		Content: text,
	}

	return c.SendMessage(channelID, body)
}
