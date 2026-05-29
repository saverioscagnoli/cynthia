package ds

type Emoji struct {
	ID            *Snowflake   `json:"id"`
	Name          *string      `json:"name"`
	Roles         *[]Snowflake `json:"roles"`
	User          *User        `json:"user"`
	RequireColons *bool        `json:"require_colons"`
	Managed       *bool        `json:"managed"`
	Animated      *bool        `json:"animated"`
	Available     *bool        `json:"available"`
}
