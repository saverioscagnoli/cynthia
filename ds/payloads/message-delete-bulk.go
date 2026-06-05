package payloads

import "cynthia/ds/dstypes"

type MessageDeleteBulk struct {
	IDs       []dstypes.Snowflake `json:"ids"`
	ChannelID dstypes.Snowflake   `json:"channel_id"`
	GuildID   *dstypes.Snowflake  `json:"guild_id"`
}
