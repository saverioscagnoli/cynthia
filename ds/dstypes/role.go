package dstypes

type Role struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	Colors       RoleColors `json:"colors"`
	Hoist        bool       `json:"hoist"`
	Icon         *string    `json:"icon"`
	UnicodeEmoji *string    `json:"unicode_emoji"`
	Position     int        `json:"position"`
	Permissions  string     `json:"permissions"`
	Managed      bool       `json:"managed"`
	Mentionable  bool       `json:"mentionable"`
	Tags         RoleTags   `json:"tags"`
	Flags        int        `json:"flags"`
}
