package dstypes

type Activity struct {
	Name              string              `json:"name"`
	Type              ActivityType        `json:"type"`
	URL               *string             `json:"url"`
	CreatedAt         int                 `json:"created_at"`
	Timestamps        *ActivityTimestamps `json:"timestamps"`
	ApplicationID     *Snowflake          `json:"application_id"`
	StatusDisplayType StatusDisplayType   `json:"status_display_type"`
	Details           *string             `json:"details"`
	DetailsURL        *string             `json:"details_url"`
	State             *string             `json:"state"`
	StateURL          *string             `json:"state_url"`
	Emoji             *Emoji              `json:"emoji"`
	Party             *ActivityParty      `json:"party"`
	Assets            *ActivityAssets     `json:"assets"`
	Secrets           *ActivitySecrets    `json:"secrets"`
	Instance          *bool               `json:"instance"`
	Flags             *ActivityFlags      `json:"flags"`
	Buttons           *[]Button           `json:"buttons"`
}
