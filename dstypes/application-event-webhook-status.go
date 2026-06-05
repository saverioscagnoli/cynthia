package dstypes

type ApplicationEventWebhookStatus int

const (
	ApplicationEventWebhookStatusDisabled          ApplicationEventWebhookStatus = 1
	ApplicationEventWebhookStatusEnabled           ApplicationEventWebhookStatus = 2
	ApplicationEventWebhookStatusDisabledByDiscord ApplicationEventWebhookStatus = 3
)
