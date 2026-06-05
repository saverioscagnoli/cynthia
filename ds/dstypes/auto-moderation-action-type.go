package dstypes

type AutoModerationActionType int

const (
	AutoModerationActionTypeBlockMessage           AutoModerationActionType = 1
	AutoModerationActionTypeSendAlertMessage       AutoModerationActionType = 2
	AutoModerationActionTypeTimeout                AutoModerationActionType = 3
	AutoModerationActionTypeBlockMemberInteraction AutoModerationActionType = 4
)
