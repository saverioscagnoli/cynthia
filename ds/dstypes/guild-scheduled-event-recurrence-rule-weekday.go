package dstypes

import "cynthia/util"

type GuildScheduledEventRecurrenceRuleWeekday int

const (
	GuildScheduledEventRecurrenceRuleWeekdayMonday    GuildScheduledEventRecurrenceRuleWeekday = 0
	GuildScheduledEventRecurrenceRuleWeekdayTuesday   GuildScheduledEventRecurrenceRuleWeekday = 1
	GuildScheduledEventRecurrenceRuleWeekdayWednesday GuildScheduledEventRecurrenceRuleWeekday = 2
	GuildScheduledEventRecurrenceRuleWeekdayThursday  GuildScheduledEventRecurrenceRuleWeekday = 3
	GuildScheduledEventRecurrenceRuleWeekdayFriday    GuildScheduledEventRecurrenceRuleWeekday = 4
	GuildScheduledEventRecurrenceRuleWeekdaySaturday  GuildScheduledEventRecurrenceRuleWeekday = 5
	GuildScheduledEventRecurrenceRuleWeekdaySunday    GuildScheduledEventRecurrenceRuleWeekday = 6
)

func (e *GuildScheduledEventRecurrenceRuleWeekday) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, e)
}
