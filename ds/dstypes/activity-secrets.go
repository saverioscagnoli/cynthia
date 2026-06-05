package dstypes

type ActivitySecrets struct {
	Join     *string `json:"join"`
	Spectate *string `json:"spectate"`
	Match    *string `json:"match"`
}
