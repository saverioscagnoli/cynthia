package dstypes

type VoiceState struct {
	GuildID                 *Snowflake   `json:"guild_id"`
	ChannelID               *Snowflake   `json:"channel_id"`
	UserID                  Snowflake    `json:"user_id"`
	Member                  *GuildMember `json:"member"`
	SessionID               string       `json:"session_id"`
	Deaf                    bool         `json:"deaf"`
	Mute                    bool         `json:"mute"`
	SelfDeaf                bool         `json:"self_deaf"`
	SelfMute                bool         `json:"self_mute"`
	SelfStream              *bool        `json:"self_stream"`
	SelfVideo               bool         `json:"self_video"`
	Suppress                bool         `json:"suppress"`
	RequestToSpeakTimestamp string       `json:"request_to_speak_timestamp"`
}
