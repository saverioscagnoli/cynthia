package dstypes

import "cynthia/util"

type RuleEventType int

const (
	RuleEventTypeMessageSend  RuleEventType = 1
	RuleEventTypeMemberUpdate RuleEventType = 2
)

func (t *RuleEventType) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, t)
}
