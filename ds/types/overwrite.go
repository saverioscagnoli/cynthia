package ds

type Overwrite struct {
	ID    Snowflake `json:"id"`
	Type  int       `json:"type"`
	Allow string    `json:"allow"`
	Deny  string    `json:"deny"`
}
