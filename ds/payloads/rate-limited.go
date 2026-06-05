package payloads

type RateLimited struct {
	Opcode     Op                                  `json:"opcode"`
	RetryAfter float32                             `json:"retry_after"`
	Meta       RequestGuildMemberRateLimitMetadata `json:"meta"`
}
