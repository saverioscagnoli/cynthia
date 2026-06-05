package dstypes

type RuleEventType = int

const (
	RuleEventTypeMessageSend  RuleEventType = 1
	RuleEventTypeMemberUpdate RuleEventType = 2
)
