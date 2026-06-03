package ds

type ThreadMetadata struct {
	Archived            bool    `json:"archived"`
	AutoArchiveDuration int     `json:"auto_archive_duration"`
	ArchiveTimestamp    string  `json:"archive_timestamp"`
	Locked              bool    `json:"locked"`
	Invitable           *bool   `json:"invitable"`
	CreateTimestamp     *string `json:"create_timestamp"`
}
