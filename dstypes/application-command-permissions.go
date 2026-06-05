package dstypes

type ApplicationCommandPermissions struct {
	ID            Snowflake                       `json:"id"`
	ApplicationID Snowflake                       `json:"application_id"`
	GuildID       Snowflake                       `json:"guild_id"`
	Permissions   []ApplicationCommandPermissions `json:"permissions"`
}
