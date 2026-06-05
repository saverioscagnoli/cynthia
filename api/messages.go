package api

import (
	"cynthia/dstypes"
	"fmt"
)

func EndpointCreateMessage(channelID dstypes.Snowflake) string {
	return fmt.Sprintf("%s/channels/%s/messages", apiLink, channelID)
}
