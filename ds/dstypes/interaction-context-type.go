package dstypes

import "cynthia/util"

type InteractionContextType int

const (
	InteractionContextTypeGuild          InteractionContextType = 0
	InteractionContextTypeBotDM          InteractionContextType = 1
	InteractionContextTypePrivateChannel InteractionContextType = 2
)

func (e *InteractionContextType) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, e)
}
