package dstypes

import "cynthia/util"

type GuildScheduledEventRecurrenceRuleFrequency int

const (
	GuildScheduledEventRecurrenceRuleFrequencyYearly  GuildScheduledEventRecurrenceRuleFrequency = 0
	GuildScheduledEventRecurrenceRuleFrequencyMonthly GuildScheduledEventRecurrenceRuleFrequency = 1
	GuildScheduledEventRecurrenceRuleFrequencyWeekly  GuildScheduledEventRecurrenceRuleFrequency = 2
	GuildScheduledEventRecurrenceRuleFrequencyDaily   GuildScheduledEventRecurrenceRuleFrequency = 3
)

func (e *GuildScheduledEventRecurrenceRuleFrequency) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, e)
}
