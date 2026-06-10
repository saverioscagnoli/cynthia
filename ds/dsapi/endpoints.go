package dsapi

import (
	"cynthia/ds/dstypes"
	"fmt"
)

func EndpointCreateMessage(channelID dstypes.Snowflake) string {
	return fmt.Sprintf("%s/channels/%s/messages", ApiURL, channelID)
}

func EndpointRegisterGuildCommand(ApplicationID dstypes.Snowflake, GuildID dstypes.Snowflake) string {
	return fmt.Sprintf("%s/applications/%s/guilds/%s/commands", ApiURL, ApplicationID, GuildID)
}
