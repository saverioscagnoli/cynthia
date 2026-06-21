package ds

import (
	"cynthia/service/util"
	"strconv"
)

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

func (u *User) AvatarURL() string {
	return Routes.AvatarURL(u.ID, *u.Avatar)
}

func (u *User) AvatarURLSize(size int) string {
	return Routes.AvatarURL(u.ID, *u.Avatar) + "?size=" + strconv.Itoa(size)
}

type AvatarDecorationData struct {
	Asset string `json:"asset"`
	SkuID string `json:"sku_id"`
}

type Collectibles struct {
	Nameplate *Nameplate `json:"nameplate"`
}

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

type EntitlementType int

const (
	EntitlementTypePurchase                EntitlementType = 1
	EntitlementTypePremiumSubscription     EntitlementType = 2
	EntitlementTypeDeveloperGift           EntitlementType = 3
	EntitlementTypeTestModePurchase        EntitlementType = 4
	EntitlementTypeFreePurchase            EntitlementType = 5
	EntitlementTypeUserGift                EntitlementType = 6
	EntitlementTypePremiumPurchase         EntitlementType = 7
	EntitlementTypeApplicationSubscription EntitlementType = 8
)

func (e *EntitlementType) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, e)
}

type Nameplate struct {
	SkuID   Snowflake `json:"sku_id"`
	Asset   string    `json:"asset"`
	Label   string    `json:"label"`
	Palette string    `json:"palette"`
}

type UserPrimaryGuild struct {
	IdentityGuildID *Snowflake `json:"identity_guild_id"`
	IdentityEnabled *bool      `json:"identity_enabled"`
	Tag             *string    `json:"tag"`
	Badge           *string    `json:"badge"`
}
