package dstypes

import "cynthia/util"

type ApplicationEventWebhookStatus int

const (
	ApplicationEventWebhookStatusDisabled          ApplicationEventWebhookStatus = 1
	ApplicationEventWebhookStatusEnabled           ApplicationEventWebhookStatus = 2
	ApplicationEventWebhookStatusDisabledByDiscord ApplicationEventWebhookStatus = 3
)

func (a *ApplicationEventWebhookStatus) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, a)
}
