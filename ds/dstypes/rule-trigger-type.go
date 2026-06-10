package dstypes

import "cynthia/util"

type RuleTriggerType int

const (
	RuleTriggerTypeKeyword       RuleTriggerType = 1
	RuleTriggerTypeSpam          RuleTriggerType = 3
	RuleTriggerTypeKeywordPreset RuleTriggerType = 4
	RuleTriggerTypeMentionSpam   RuleTriggerType = 5
	RuleTriggerTypeMemberProfile RuleTriggerType = 6
)

func (t *RuleTriggerType) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, t)
}
