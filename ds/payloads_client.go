package ds

type Ready struct {
	Version int  `json:"v"`
	User    User `json:"user"`
}
