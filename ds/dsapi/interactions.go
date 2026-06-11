package dsapi

import (
	"cynthia/ds/dstypes"
	"fmt"
)

func EnpointInteractionCallback(intID dstypes.Snowflake, intToken string) string {
	return fmt.Sprintf("%s/interactions/%s/%s/callback", ApiURL, intID, intToken)
}

func EndpointFollowupMessage(appID dstypes.Snowflake, token string) string {
	return fmt.Sprintf("%s/webhooks/%s/%s", ApiURL, appID, token)
}

func (c *Client) InteractionCallback(interaction *dstypes.Interaction, response dstypes.InteractionResponse) error {
	_, err := Request[any](c, "POST", EnpointInteractionCallback(interaction.ID, interaction.Token), response)
	return err
}

func (c *Client) InteractionReply(interaction *dstypes.Interaction, data dstypes.InteractionCallbackData) error {
	return c.InteractionCallback(interaction, dstypes.InteractionResponse{
		Type: dstypes.InteractionCallbackTypeChannelMessageWithSource,
		Data: &data,
	})
}

func (c *Client) InteractionReplyText(interaction *dstypes.Interaction, message string) error {
	return c.InteractionReply(interaction, dstypes.InteractionCallbackData{Content: message})
}

func (c *Client) InteractionDefer(interaction *dstypes.Interaction) error {
	return c.InteractionCallback(interaction, dstypes.InteractionResponse{
		Type: dstypes.InteractionCallbackTypeDeferredChannelMessageWithSource,
	})
}

func (c *Client) InteractionFollowup(interaction *dstypes.Interaction, data dstypes.InteractionCallbackData) error {
	_, err := Request[any](c, "POST", EndpointFollowupMessage(c.appID, interaction.Token), data)
	return err
}

func (c *Client) InteractionFollowupText(interaction *dstypes.Interaction, message string) error {
	return c.InteractionFollowup(interaction, dstypes.InteractionCallbackData{Content: message})
}
