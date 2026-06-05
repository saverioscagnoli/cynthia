package payloads

import "cynthia/dstypes"

type GuildSoundboardSoundDelete struct {
	SoundID dstypes.Snowflake `json:"sound_id"`
	GuildID dstypes.Snowflake `json:"guild_id"`
}
