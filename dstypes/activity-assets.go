package dstypes

type ActivityAssets struct {
	LargeImage       *string `json:"large_image"`
	LargeText        *string `json:"large_text"`
	LargeURL         *string `json:"large_url"`
	SmallImage       *string `json:"small_image"`
	SmallText        *string `json:"small_text"`
	SmallURL         *string `json:"small_url"`
	InviteCoverImage *string `json:"invite_cover_image"`
}
