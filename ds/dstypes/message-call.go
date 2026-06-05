package dstypes

type MessageCall struct {
	Participants   []Snowflake `json:"participants"`
	EndedTimestamp *string     `json:"ended_timestamp"`
}
