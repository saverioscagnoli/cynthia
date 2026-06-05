package dstypes

type TriggerMetadata struct {
	KeywordFilter                []string             `json:"keyword_filter"`
	RegexPatterns                []string             `json:"regex_patterns"`
	Presets                      []KeywordPresetTypes `json:"presets"`
	AllowList                    []string             `json:"allow_list"`
	MentionTotalLimit            int                  `json:"mention_total_limit"`
	MentionRaidProtectionEnabled bool                 `json:"mention_raid_protection_enabled"`
}
