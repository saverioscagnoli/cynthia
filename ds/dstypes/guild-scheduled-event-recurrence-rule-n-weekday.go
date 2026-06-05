package dstypes

type GuildScheduledEventRecurrenceRuleNWeekday struct {
	N   int                                      `json:"n"`
	Day GuildScheduledEventRecurrenceRuleWeekday `json:"day"`
}
