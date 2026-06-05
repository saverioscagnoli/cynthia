package payloads

import "cynthia/dstypes"

type SoundboardSounds struct {
	SoundboardSounds []dstypes.SoundboardSound `json:"soundboard_sounds"`
	GuildID          dstypes.Snowflake         `json:"guild_id"`
}
