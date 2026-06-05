package payloads

import "cynthia/dstypes"

type GuildEmojisUpdate struct {
	GuildID dstypes.Snowflake `json:"guild_id"`
	Emojis  []dstypes.Emoji   `json:"emojis"`
}
