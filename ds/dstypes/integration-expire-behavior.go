package dstypes

import "cynthia/util"

type IntegrationExpireBehavior int

const (
	IntegrationExpireBehaviorRemoveRole IntegrationExpireBehavior = 0
	IntegrationExpireBehaviorKick       IntegrationExpireBehavior = 1
)

func (e *IntegrationExpireBehavior) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, e)
}
