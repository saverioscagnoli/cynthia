package dstypes

type InteractionResponse struct {
	Type InteractionCallbackType  `json:"type"`
	Data *InteractionCallbackData `json:"data"`
}
