package payloads

import "cynthia/dstypes"

type InviteDelete struct {
	ChannelID dstypes.Snowflake  `json:"channel_id"`
	GuildID   *dstypes.Snowflake `json:"guild_id"`
	Code      string             `json:"code"`
}
