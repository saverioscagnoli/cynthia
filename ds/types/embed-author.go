package ds

type EmbedAuthor struct {
	Name         *string `json:"name"`
	URL          *string `json:"url"`
	IconURL      *string `json:"icon_url"`
	ProxyIconURL *string `json:"proxy_icon_url"`
}
