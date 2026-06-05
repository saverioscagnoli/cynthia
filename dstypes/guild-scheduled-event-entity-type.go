package dstypes

type GuildScheduledEventEntityType int

const (
	GuildScheduledEventEntityTypeStageInstace GuildScheduledEventEntityType = 1
	GuildScheduledEventEntityTypeVoice        GuildScheduledEventEntityType = 2
	GuildScheduledEventEntityTypeExternal     GuildScheduledEventEntityType = 3
)
