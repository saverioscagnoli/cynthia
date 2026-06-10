package dstypes

import "cynthia/util"

type GuildScheduledEventEntityType int

const (
	GuildScheduledEventEntityTypeStageInstace GuildScheduledEventEntityType = 1
	GuildScheduledEventEntityTypeVoice        GuildScheduledEventEntityType = 2
	GuildScheduledEventEntityTypeExternal     GuildScheduledEventEntityType = 3
)

func (e *GuildScheduledEventEntityType) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, e)
}
