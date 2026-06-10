package dstypes

import "cynthia/util"

type EntryPointCommandHandlerType int

const (
	EntryPointCommandHandlerTypeAppHandler            EntryPointCommandHandlerType = 0
	EntryPointCommandHandlerTypeDiscordLaunchActivity EntryPointCommandHandlerType = 1
)

func (e *EntryPointCommandHandlerType) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, e)
}
