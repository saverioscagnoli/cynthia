package dstypes

import "cynthia/util"

type ApplicationIntegrationType int

const (
	ApplicationIntegrationTypesGuildInstall ApplicationIntegrationType = 0
	ApplicationIntegrationTypesUserInstall  ApplicationIntegrationType = 1
)

func (a *ApplicationIntegrationType) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, a)
}
