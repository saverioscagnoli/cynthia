package dstypes

type Interaction struct {
	ID     Snowflake
	Type   InteractionType
	Name   string
	User   User
	Member GuildMember
}
