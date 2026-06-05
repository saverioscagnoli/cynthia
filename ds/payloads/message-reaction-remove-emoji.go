package payloads

import "cynthia/ds/dstypes"

type MessageReactionRemoveEmoji struct {
	ChannelID dstypes.Snowflake  `json:"channel_id"`
	GuildID   *dstypes.Snowflake `json:"guild_id"`
	MessageID dstypes.Snowflake  `json:"message_id"`
	Emoji     dstypes.Emoji      `json:"emoji"`
}
