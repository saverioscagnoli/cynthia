package dstypes

import "cynthia/util"

type MembershipState int

const (
	MembershipStateInvited  = 1
	MembershipStateAccepted = 2
)

func (t *MembershipState) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, t)
}

type TeamMember struct {
	MembershipState MembershipState `json:"membership_state"`
	TeamID          Snowflake       `json:"team_id"`
	User            User            `json:"user"`
	Role            string          `json:"role"`
}
