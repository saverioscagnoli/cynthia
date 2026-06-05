package payloads

import "cynthia/dstypes"

type ThreadMembersUpdate struct {
	ID               dstypes.Snowflake       `json:"id"`
	GuildID          dstypes.Snowflake       `json:"guild_id"`
	MemberCount      int                     `json:"member_count"`
	AddedMembers     *[]dstypes.ThreadMember `json:"added_members"`
	RemovedMemberIDs *[]dstypes.Snowflake    `json:"removed_member_ids"`
}
