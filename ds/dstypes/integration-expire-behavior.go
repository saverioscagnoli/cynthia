package dstypes

type IntegrationExpireBehavior int

const (
	IntegrationExpireBehaviorRemoveRole IntegrationExpireBehavior = 0
	IntegrationExpireBehaviorKick       IntegrationExpireBehavior = 1
)
