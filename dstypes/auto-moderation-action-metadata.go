package dstypes

type AutoModerationActionMetadata struct {
	ChannelID       Snowflake `json:"channel_id"`
	DurationSeconds int       `json:"duration_seconds"`
	CustomMessage   *string   `json:"custom_message"`
}
