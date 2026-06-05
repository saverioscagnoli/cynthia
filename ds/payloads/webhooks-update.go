package payloads

import "cynthia/dstypes"

type WebhooksUpdate struct {
	GuildID   dstypes.Snowflake `json:"guild_id"`
	ChannelID dstypes.Snowflake `json:"channel_id"`
}
