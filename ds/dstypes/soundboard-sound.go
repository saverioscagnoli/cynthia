package dstypes

type SoundboardSound struct {
	Name      string     `json:"name"`
	SoundID   Snowflake  `json:"sound_id"`
	Volume    float64    `json:"volume"`
	EmojiID   *Snowflake `json:"emoji_id"`
	EmojiName *string    `json:"emoji_name"`
	GuildID   *Snowflake `json:"guild_id"`
	Available bool       `json:"available"`
	User      *User      `json:"user"`
}
