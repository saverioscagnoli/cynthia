package dstypes

import "cynthia/util"

type GuildScheduledEventRecurrenceRuleMonth int

const (
	GuildScheduledEventRecurrenceRuleMonthJanuary   GuildScheduledEventRecurrenceRuleMonth = 1
	GuildScheduledEventRecurrenceRuleMonthFebruary  GuildScheduledEventRecurrenceRuleMonth = 2
	GuildScheduledEventRecurrenceRuleMonthMarch     GuildScheduledEventRecurrenceRuleMonth = 3
	GuildScheduledEventRecurrenceRuleMonthApril     GuildScheduledEventRecurrenceRuleMonth = 4
	GuildScheduledEventRecurrenceRuleMonthMay       GuildScheduledEventRecurrenceRuleMonth = 5
	GuildScheduledEventRecurrenceRuleMonthJune      GuildScheduledEventRecurrenceRuleMonth = 6
	GuildScheduledEventRecurrenceRuleMonthJuly      GuildScheduledEventRecurrenceRuleMonth = 7
	GuildScheduledEventRecurrenceRuleMonthAugust    GuildScheduledEventRecurrenceRuleMonth = 8
	GuildScheduledEventRecurrenceRuleMonthSeptember GuildScheduledEventRecurrenceRuleMonth = 9
	GuildScheduledEventRecurrenceRuleMonthOctober   GuildScheduledEventRecurrenceRuleMonth = 10
	GuildScheduledEventRecurrenceRuleMonthNovember  GuildScheduledEventRecurrenceRuleMonth = 11
	GuildScheduledEventRecurrenceRuleMonthDecember  GuildScheduledEventRecurrenceRuleMonth = 12
)

func (e *GuildScheduledEventRecurrenceRuleMonth) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, e)
}
