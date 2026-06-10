package dsapi

import (
	"cynthia/ds/dstypes"
	"fmt"
)

func EnpointInteractionCalback(intID dstypes.Snowflake, intToken string) string {
	return fmt.Sprintf("%s/interactions/%s/%s/callback", ApiURL, intID, intToken)
}

func (c *Client) InteractionReply(intID dstypes.Snowflake, intToken string, data dstypes.InteractionResponse) error {
	_, err := c.Request("POST", EnpointInteractionCalback(intID, intToken), data)
	return err
}
