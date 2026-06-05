package payloads

import "cynthia/dstypes"

type VoiceChannelEffectSend struct {
	ChannelID     dstypes.Snowflake      `json:"channel_id"`
	GuildID       dstypes.Snowflake      `json:"guild_id"`
	UserID        dstypes.Snowflake      `json:"user_id"`
	Emoji         *dstypes.Emoji         `json:"emoji"`
	AnimationType *dstypes.AnimationType `json:"animation_type"`
	AnimationID   *int                   `json:"animation_id"`
	SoundID       *dstypes.Snowflake     `json:"sound_id"`
	SoundVolume   float64                `json:"sound_volume"`
}
