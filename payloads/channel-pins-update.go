package payloads

import "cynthia/dstypes"

type ChannelPinsUpdate struct {
	GuildID          *dstypes.Snowflake `json:"guild_id"`
	ChannelID        dstypes.Snowflake  `json:"channel_id"`
	LastPinTimestamp *string            `json:"last_pin_timestamp"`
}
