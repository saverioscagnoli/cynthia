package ds

type MembershipState int

const (
	MembershipStateInvited  = 1
	MembershipStateAccepted = 2
)

type TeamMember struct {
	MembershipState MembershipState `json:"membership_state"`
	TeamID          Snowflake       `json:"team_id"`
	User            User            `json:"user"`
	Role            string          `json:"role"`
}
