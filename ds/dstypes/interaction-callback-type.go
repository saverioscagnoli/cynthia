package dstypes

import "cynthia/util"

type InteractionCallbackType int

const (
	InteractionCallbackTypePong                                 InteractionCallbackType = 1
	InteractionCallbackTypeChannelMessageWithSource             InteractionCallbackType = 4
	InteractionCallbackTypeDeferredChannelMessageWithSource     InteractionCallbackType = 5
	InteractionCallbackTypeDeferredUpdateMessage                InteractionCallbackType = 6
	InteractionCallbackTypeUpdateMessage                        InteractionCallbackType = 7
	InteractionCallbackTypeApplicationCommandAutoCompleteResult InteractionCallbackType = 8
	InteractionCallbackTypeModal                                InteractionCallbackType = 9
	InteractionCallbackTypePremiumRequired                      InteractionCallbackType = 10
	InteractionCallbackTypeLaunchActivity                       InteractionCallbackType = 11
)

func (t *InteractionCallbackType) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, t)
}
