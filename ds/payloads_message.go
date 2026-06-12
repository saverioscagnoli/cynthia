package ds

type MessageCreate struct {
	Message
	GuildID     *Snowflake   `json:"guild_id"`
	Member      *GuildMember `json:"member"`
	Mentions    *[]User      `json:"mentions"`
	ChannelType *ChannelType `json:"channel_type"`
}
