package dstypes

type UserPrimaryGuild struct {
	IdentityGuildID *Snowflake `json:"identity_guild_id"`
	IdentityEnabled *bool      `json:"identity_enabled"`
	Tag             *string    `json:"tag"`
	Badge           *string    `json:"badge"`
}
