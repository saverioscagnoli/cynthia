package ds

type GuildMember struct {
	User                       *User                 `json:"user"`
	Nick                       *string               `json:"nick"`
	Avatar                     *string               `json:"avatar"`
	Banner                     *string               `json:"banner"`
	Roles                      []Snowflake           `json:"roles"`
	JoinedAt                   string                `json:"joined_at"`
	PremiumSince               *string               `json:"premium_since"`
	Deaf                       bool                  `json:"deaf"`
	Mute                       bool                  `json:"mute"`
	Flags                      int                   `json:"flags"`
	Pending                    *bool                 `json:"pending"`
	Permissions                *string               `json:"permissions"`
	CommunicationDisabledUntil *string               `json:"communication_disabled_until"`
	AvatarDecorationData       *AvatarDecorationData `json:"avatar_decoration_data"`
	Collectibles               *Collectibles         `json:"collectibles"`
}
