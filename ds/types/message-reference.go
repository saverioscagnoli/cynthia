package ds

type MessageReference struct {
	Type            *int       `json:"type"`
	MessageID       *Snowflake `json:"message_id"`
	ChannelID       *Snowflake `json:"channel_id"`
	GuildID         *Snowflake `json:"guild_id"`
	FailIfNotExists *bool      `json:"fail_if_not_exists"`
}
