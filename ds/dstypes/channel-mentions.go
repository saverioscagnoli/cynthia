package dstypes

type ChannelMentions struct {
	ID      Snowflake `json:"id"`
	GuildID Snowflake `json:"guild_id"`
	Type    int       `json:"type"`
	Name    string    `json:"name"`
}
