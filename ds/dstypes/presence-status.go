package dstypes

type PresenceStatus string

const (
	PresenceStatusIdle    PresenceStatus = "idle"
	PresenceStatusDnd     PresenceStatus = "dnd"
	PresenceStatusOnline  PresenceStatus = "online"
	PresenceStatusOffline PresenceStatus = "offline"
)
