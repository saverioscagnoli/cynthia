package payloads

import "cynthia/ds/dstypes"

type GuildStickersUpdate struct {
	GuildID  dstypes.Snowflake `json:"guild_id"`
	Stickers []dstypes.Sticker `json:"stickers"`
}
