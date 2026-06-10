package dstypes

import "cynthia/util"

type PresenceStatus string

const (
	PresenceStatusIdle    PresenceStatus = "idle"
	PresenceStatusDnd     PresenceStatus = "dnd"
	PresenceStatusOnline  PresenceStatus = "online"
	PresenceStatusOffline PresenceStatus = "offline"
)

func (s *PresenceStatus) UnmarshalJSON(data []byte) error {
	return util.UnmarshalString(data, s)
}
