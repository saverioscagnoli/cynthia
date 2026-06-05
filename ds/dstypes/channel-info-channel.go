package dstypes

type ChannelInfoChannel struct {
	ID             Snowflake `json:"id"`
	Status         *string   `json:"status"`
	VoiceStartTime *int      `json:"voice_start_time"`
}
