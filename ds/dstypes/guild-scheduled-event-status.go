package dstypes

import "cynthia/util"

type GuildScheduledEventStatus int

const (
	GuildScheduledEventStatusScheduled GuildScheduledEventStatus = 1
	GuildScheduledEventStatusActive    GuildScheduledEventStatus = 2
	GuildScheduledEventCompleted       GuildScheduledEventStatus = 3
	GuildScheduledEventCanceled        GuildScheduledEventStatus = 4
)

func (e *GuildScheduledEventStatus) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, e)
}
