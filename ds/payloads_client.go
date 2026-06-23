package ds

type Ready struct {
	Version int  `json:"v"`
	User    User `json:"user"`
}

type Resumed struct{}

type Reconnect struct{}

type RateLimited struct {
	Opcode     Op                                  `json:"opcode"`
	RetryAfter float32                             `json:"redtry_after"`
	Meta       RequestGuildMemberRateLimitMetadata `json:"meta"`
}

type InvalidSession struct {
	bool
}
