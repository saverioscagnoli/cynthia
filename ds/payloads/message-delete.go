package payloads

import "cynthia/ds/dstypes"

type MessageDelete struct {
	ID        dstypes.Snowflake  `json:"id"`
	ChannelID dstypes.Snowflake  `json:"channel_id"`
	GuildID   *dstypes.Snowflake `json:"guild_id"`
}
