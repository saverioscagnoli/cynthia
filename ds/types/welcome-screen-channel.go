package ds

type WelcomeScreenChannel struct {
	ChannelID   Snowflake  `json:"channel_id"`
	Description string     `json:"description"`
	EmojiID     *Snowflake `json:"emoji_id"`
	EmojiName   *string    `json:"emoji_name"`
}
