package dstypes

import "cynthia/util"

type ReactionType int

const (
	ReactionTypeNormal ReactionType = 0
	ReactionTypeBurst  ReactionType = 1
)

func (t *ReactionType) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, t)
}
