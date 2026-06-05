package dstypes

type GuildScheduledEventStatus int

const (
	GuildScheduledEventStatusScheduled GuildScheduledEventStatus = 1
	GuildScheduledEventStatusActive    GuildScheduledEventStatus = 2
	GuildScheduledEventCompleted       GuildScheduledEventStatus = 3
	GuildScheduledEventCanceled        GuildScheduledEventStatus = 4
)
