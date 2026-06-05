package dstypes

type RuleTriggerType int

const (
	RuleTriggerTypeKeyword       RuleTriggerType = 1
	RuleTriggerTypeSpam          RuleTriggerType = 3
	RuleTriggerTypeKeywordPreset RuleTriggerType = 4
	RuleTriggerTypeMentionSpam   RuleTriggerType = 5
	RuleTriggerTypeMemberProfile RuleTriggerType = 6
)
