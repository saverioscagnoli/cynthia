package models

import "cynthia/ds"

type User struct {
	DiscordID ds.Snowflake `json:"discord_id"`
	Username  string       `json:"username"`
	Avatar    *string      `json:"avatar"`
	TrainerID int          `json:"trainer_id"`
	CreatedAt string       `json:"created_at"`
}
