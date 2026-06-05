package dstypes

type Nameplate struct {
	SkuID   Snowflake `json:"sku_id"`
	Asset   string    `json:"asset"`
	Label   string    `json:"label"`
	Palette string    `json:"palette"`
}
