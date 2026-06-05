package payloads

import "cynthia/dstypes"

type MessagePollVoteRemove struct {
	UserID    dstypes.Snowflake  `json:"user_id"`
	ChannelID dstypes.Snowflake  `json:"channel_id"`
	MessageID dstypes.Snowflake  `json:"message_id"`
	GuildID   *dstypes.Snowflake `json:"guild_id"`
	AnswerID  int
}
