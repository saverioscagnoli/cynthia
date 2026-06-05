package dstypes

type IncidentsData struct {
	InvitesDisabledUntil string `json:"invites_disabled_until"`
	DmsDisabledUntil     string `json:"dms_disabled_until"`
	DmSpamDetectedAt     string `json:"dm_spam_detected_at"`
	RaidDetectedAt       string `json:"raid_detected_at"`
}
