package api

import (
	"cynthia/ds/dstypes"
	"fmt"
)

func EndpointCreateMessage(channelID dstypes.Snowflake) string {
	return fmt.Sprintf("%s/channels/%s/messages", apiLink, channelID)
}
