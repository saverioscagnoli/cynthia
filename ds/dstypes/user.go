package dstypes

type User struct {
	ID                   Snowflake             `json:"id"`
	Username             string                `json:"username"`
	Discriminator        string                `json:"discriminator"`
	GlobalName           *string               `json:"global_name"`
	Avatar               *string               `json:"avatar"`
	Bot                  *bool                 `json:"bot"`
	System               *bool                 `json:"system"`
	Mfa                  *bool                 `json:"mfa"`
	Banner               *string               `json:"banner"`
	AccentColor          *int                  `json:"accent_color"`
	Locale               *string               `json:"locale"`
	Verified             *bool                 `json:"verified"`
	Email                *string               `json:"email"`
	Flags                *int                  `json:"flags"`
	PremiumType          *int                  `json:"premium_type"`
	PublicFlags          *int                  `json:"public_flags"`
	AvatarDecorationData *AvatarDecorationData `json:"avatar_decoration_data"`
	Collectibles         *Collectibles         `json:"collectibles"`
	PrimaryGuild         *UserPrimaryGuild     `json:"primary_guild"`
}
