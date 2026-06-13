package ds

import (
	"cynthia/util"
	"fmt"
	"net/http"
)

func (r *routes) InteractionCallback(intID Snowflake, intToken string) (string, string) {
	return http.MethodPost, fmt.Sprintf("/interactions/%s/%s/callback", intID, intToken)
}

func (r *routes) InteractionFollowup(appID Snowflake, intToken Snowflake) (string, string) {
	return http.MethodPost, fmt.Sprintf("/webhooks/%s/%s", appID, intToken)
}

func (c *ApiClient) InteractionCallback(i *Interaction, data InteractionResponse) error {
	method, endpoint := Routes.InteractionCallback(i.ID, i.Token)
	_, err := c.request(method, endpoint, data)

	return err
}

func (c *ApiClient) InteractionReply(i *Interaction, data *InteractionCallbackData) error {
	return c.InteractionCallback(i, InteractionResponse{
		Type: InteractionCallbackTypeChannelMessageWithSource,
		Data: data,
	})
}

func (c *ApiClient) InteractionReplyText(i *Interaction, message string) error {
	return c.InteractionReply(i, &InteractionCallbackData{Content: message})
}

func (c *ApiClient) InteractionDefer(i *Interaction) error {
	return c.InteractionCallback(i, InteractionResponse{
		Type: InteractionCallbackTypeDeferredChannelMessageWithSource,
	})
}

func (c *ApiClient) InteractionFollowup(i *Interaction, data InteractionCallbackData) error {
	method, endpoint := Routes.InteractionFollowup(c.appID, i.Token)
	_, err := c.request(method, endpoint, data)
	return err
}

func (c *ApiClient) InteractionFollowupText(i *Interaction, message string) error {
	return c.InteractionFollowup(i, InteractionCallbackData{Content: message})
}

func (c *ApiClient) InteractionReplyEphemeral(i *Interaction, data *InteractionCallbackData) error {
	data.Flags = util.Ptr(64)
	return c.InteractionReply(i, data)
}

func (c *ApiClient) InteractionReplyTextEphemeral(i *Interaction, message string) error {
	return c.InteractionReplyEphemeral(i, &InteractionCallbackData{Content: message})
}
