package payloads

import "cynthia/dstypes"

type VoiceChannelStartTimeUpdate struct {
	ID             dstypes.Snowflake `json:"id"`
	GuildID        dstypes.Snowflake `json:"guild_id"`
	VoiceStartTime *int              `json:"voice_start_time"`
}
