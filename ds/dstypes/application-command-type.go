package dstypes

import "cynthia/util"

type ApplicationCommandType int

const (
	ApplicationCommandTypeChatInput         ApplicationCommandType = 1
	ApplicationCommandTypeUser              ApplicationCommandType = 2
	ApplicationCommandTypeMessage           ApplicationCommandType = 3
	ApplicationCommandTypePrimaryEntryPoint ApplicationCommandType = 4
)

func (a *ApplicationCommandType) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, a)
}
