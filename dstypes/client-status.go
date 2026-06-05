package dstypes

type ClientStatus struct {
	Desktop *string `json:"desktop"`
	Mobile  *string `json:"mobile"`
	Web     *string `json:"web"`
}
