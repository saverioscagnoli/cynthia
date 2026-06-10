package dstypes

import "cynthia/util"

type AnimationType int

const (
	AnimationTypePremium AnimationType = 0
	AnimationTypeBasic   AnimationType = 1
)

func (a *AnimationType) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, a)
}
