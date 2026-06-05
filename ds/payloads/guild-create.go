package payloads

import (
	"cynthia/ds/dstypes"
)

type GuildCreate struct {
	dstypes.Guild
	JoinedAt             string                        `json:"joined_at"`
	Unavailable          *bool                         `json:"unavailable"`
	MemberCount          int                           `json:"member_count"`
	VoiceStates          []dstypes.VoiceState          `json:"voice_states"`
	Members              []dstypes.GuildMember         `json:"members"`
	Channels             []dstypes.Channel             `json:"channels"`
	Threads              []dstypes.Channel             `json:"threads"`
	Presences            []PresenceUpdate              `json:"presences"`
	StageInstances       []dstypes.StageInstance       `json:"stage_instances"`
	GuildScheduledEvents []dstypes.GuildScheduledEvent `json:"guild_scheduled_events"`
	SoundboardSounts     []dstypes.SoundboardSound     `json:"soundboard_sounds"`
}
