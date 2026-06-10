package dstypes

import "cynthia/util"

type AllowedMentions string

const (
	AllowedMentionsRole     AllowedMentions = "role"
	AllowedMentionsUser     AllowedMentions = "user"
	AllowedMentionsEveryone AllowedMentions = "everyone"
)

func (a *AllowedMentions) UnmarshalJSON(b []byte) error {
	return util.UnmarshalString(b, a)
}
