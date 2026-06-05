package dstypes

type Entitlement struct {
	ID            Snowflake       `json:"id"`
	SkuID         Snowflake       `json:"sku_id"`
	ApplicationID Snowflake       `json:"application_id"`
	UserID        *Snowflake      `json:"user_id"`
	Type          EntitlementType `json:"type"`
	Deleted       bool            `json:"deleted"`
	StartsAt      *string         `json:"starts_at"`
	EndsAt        *string         `json:"ends_at"`
	GuildID       *Snowflake      `json:"guild_id"`
	Consumed      *bool           `json:"consumed"`
}
