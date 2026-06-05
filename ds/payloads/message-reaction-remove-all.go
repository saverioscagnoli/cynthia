package payloads

import "cynthia/dstypes"

type MessageReactionRemoveAll struct {
	ChannelID dstypes.Snowflake  `json:"channel_id"`
	MessageID dstypes.Snowflake  `json:"message_id"`
	GuildID   *dstypes.Snowflake `json:"guild_id"`
}
