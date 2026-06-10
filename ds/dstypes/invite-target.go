package dstypes

import "cynthia/util"

type InviteTarget int

const (
	InviteTargetStream              InviteTarget = 1
	InviteTargetEmbeddedApplication InviteTarget = 2
)

func (i *InviteTarget) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, i)
}
