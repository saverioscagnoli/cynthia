package dstypes

import "cynthia/util"

type GuildScheduledEventPrivacyLevel int

const (
	GuildScheduledEventPrivacyLevelGuildOnly GuildScheduledEventPrivacyLevel = 2
)

func (e *GuildScheduledEventPrivacyLevel) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, e)
}
