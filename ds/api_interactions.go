package ds

import (
	"cynthia/service/util"
	"fmt"
	"net/http"
)

func (r *routes) InteractionCallback(intID Snowflake, intToken string) (string, string) {
	return http.MethodPost, fmt.Sprintf("/interactions/%s/%s/callback", intID, intToken)
}

func (r *routes) InteractionFollowup(appID Snowflake, intToken string) (string, string) {
	return http.MethodPost, fmt.Sprintf("/webhooks/%s/%s", appID, intToken)
}

func (c *ApiClient) interactionCallbackMultipart(i *Interaction, data InteractionResponse) error {
	method, endpoint := Routes.InteractionCallback(
		i.ID,
		i.Token,
	)

	_, err := c.multipartRequest(
		method,
		endpoint,
		data,
		data.Data.Files,
		true,
	)

	return err
}

func (c *ApiClient) InteractionCallback(i *Interaction, res InteractionResponse,
) error {
	if res.Data != nil && len(res.Data.Files) > 0 {
		return c.interactionCallbackMultipart(i, res)
	}

	method, endpoint := Routes.InteractionCallback(i.ID, i.Token)
	_, err := c.request(method, endpoint, res)
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

func (c *ApiClient) InteractionDeferUpdate(i *InteractionCreate) error {
	return c.InteractionCallback(i, InteractionResponse{
		Type: InteractionCallbackTypeDeferredUpdateMessage,
	})
}

func (c *ApiClient) interactionFollowupMultipart(i *Interaction, data *InteractionCallbackData,
) error {
	method, endpoint := Routes.InteractionFollowup(
		c.appID,
		i.Token,
	)

	_, err := c.multipartRequest(
		method,
		endpoint,
		data,
		data.Files,
		true,
	)

	return err
}

func (c *ApiClient) InteractionFollowup(i *Interaction, data *InteractionCallbackData) error {
	if data != nil && len(data.Files) > 0 {
		return c.interactionFollowupMultipart(i, data)
	}

	method, endpoint := Routes.InteractionFollowup(
		c.appID,
		i.Token,
	)

	_, err := c.request(method, endpoint, data)

	return err
}

func (c *ApiClient) InteractionFollowupText(i *Interaction, message string) error {
	return c.InteractionFollowup(i, &InteractionCallbackData{Content: message})
}

func (c *ApiClient) InteractionReplyEphemeral(i *Interaction, data *InteractionCallbackData) error {
	data.Flags = util.Ptr(64)
	return c.InteractionReply(i, data)
}

func (c *ApiClient) InteractionReplyTextEphemeral(i *Interaction, message string) error {
	return c.InteractionReplyEphemeral(i, &InteractionCallbackData{Content: message})
}
