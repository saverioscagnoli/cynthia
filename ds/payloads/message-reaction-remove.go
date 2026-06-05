package payloads

import "cynthia/dstypes"

type MessageReactionRemove struct {
	UserID    dstypes.Snowflake    `json:"user_id"`
	ChannelID dstypes.Snowflake    `json:"channel_id"`
	MessageID dstypes.Snowflake    `json:"message_id"`
	GuildID   *dstypes.Snowflake   `json:"guild_id"`
	Emoji     *dstypes.Emoji       `json:"emoji"`
	Burst     bool                 `json:"burst"`
	Type      dstypes.ReactionType `json:"type"`
}
