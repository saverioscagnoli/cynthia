package ds

type MessageActivity struct {
	Type    int        `json:"type"`
	PartyID *Snowflake `json:"party_id"`
}
