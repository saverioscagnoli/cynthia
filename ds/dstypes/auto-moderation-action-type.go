package dstypes

import "cynthia/util"

type AutoModerationActionType int

const (
	AutoModerationActionTypeBlockMessage           AutoModerationActionType = 1
	AutoModerationActionTypeSendAlertMessage       AutoModerationActionType = 2
	AutoModerationActionTypeTimeout                AutoModerationActionType = 3
	AutoModerationActionTypeBlockMemberInteraction AutoModerationActionType = 4
)

func (a *AutoModerationActionType) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, a)
}
