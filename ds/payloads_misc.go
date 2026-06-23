package ds

type SoundboardSounds struct {
	SoundboardSounds []SoundboardSound `json:"soundboard_sounds"`
	GuildID          Snowflake         `json:"guild_id"`
}

type StageInstanceCreate = StageInstance

type StageInstanceUpdate = StageInstance

type StageInstanceDelete = StageInstance

type SubscriptionCreate = Subscription

type SubscriptionUpdate = Subscription

type SubscriptionDelete = Subscription

type VoiceStateUpdate = VoiceState

type VoiceServerUpdate struct {
	Token    string    `json:"token"`
	GuildID  Snowflake `json:"guild_id"`
	Endpoint *string   `json:"endpoint"`
}

type WebhooksUpdate struct {
	GuildID   Snowflake `json:"guild_id"`
	ChannelID Snowflake `json:"channel_id"`
}
