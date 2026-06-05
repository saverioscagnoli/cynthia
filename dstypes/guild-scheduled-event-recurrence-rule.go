package dstypes

type GuildScheduledEventRecurrenceRule struct {
	Start      string                                       `json:"start"`
	End        string                                       `json:"end"`
	Frequency  GuildScheduledEventRecurrenceRuleFrequency   `json:"frequency"`
	Interval   int                                          `json:"interval"`
	ByWeekday  *[]GuildScheduledEventRecurrenceRuleWeekday  `json:"by_weekday"`
	ByNWeekday *[]GuildScheduledEventRecurrenceRuleNWeekday `json:"by_n_weekday"`
	ByMonth    *[]GuildScheduledEventRecurrenceRuleMonth    `json:"by_month"`
	ByMonthDay *[]int                                       `json:"by_month_day"`
	ByYearDay  *[]int                                       `json:"by_year_day"`
	Count      *int                                         `json:"count"`
}
