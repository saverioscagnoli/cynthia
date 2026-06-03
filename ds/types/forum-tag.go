package ds

type ForumTag struct {
	ID        Snowflake  `json:"id"`
	Name      string     `json:"name"`
	Moderated bool       `json:"moderated"`
	EmojiID   *Snowflake `json:"emoji_id"`
	EmojiName *string    `json:"emoji_name"`
}
