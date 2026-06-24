package models

import (
	"cynthia/ds"
	"time"
)

type User struct {
	ID              ds.Snowflake `json:"id"`
	Username        string       `json:"username"`
	DiscordUsername string       `json:"discord_username"`
	AvatarHash      *string      `json:"avatar_hash"`
	Money           int          `json:"money"`
	SpriteID        *int         `json:"sprite_id"`
	Bag             []BagItem    `json:"bag"`
	Banner          *[]byte      `json:"banner"`
	CreatedAt       time.Time    `json:"created_at"`
}
