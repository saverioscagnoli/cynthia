package payloads

import "cynthia/ds/dstypes"

type VoiceServerUpdate struct {
	Token    string            `json:"token"`
	GuildID  dstypes.Snowflake `json:"guild_id"`
	Endpoint *string           `json:"endpoint"`
}
