package payloads

import "cynthia/dstypes"

type TypingStart struct {
	ChannelID dstypes.Snowflake    `json:"channel_id"`
	GuildID   *dstypes.Snowflake   `json:"guild_id"`
	UserID    dstypes.Snowflake    `json:"user_id"`
	Timestamp int                  `json:"timestamp"`
	Member    *dstypes.GuildMember `json:"member"`
}
