package dstypes

type ResolvedData struct {
	Users      *map[Snowflake]User
	Members    *map[Snowflake]GuildMember
	Roles      *map[Snowflake]Role
	Channels   *map[Snowflake]Channel
	Messages   *map[Snowflake]Message
	Attachment *map[Snowflake]Attachment
}
