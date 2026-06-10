package dstypes

import "cynthia/util"

type InteractionType int

const (
	InteractionTypePing                           InteractionType = 1
	InteractionTypeApplicationCommand             InteractionType = 2
	InteractionTypeMessageComponent               InteractionType = 3
	InteractionTypeApplicationCommandAutocomplete InteractionType = 4
	InteractionTypeModalSubmit                    InteractionType = 5
)

func (i *InteractionType) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, i)
}
