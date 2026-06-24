package models

type BagItem struct {
	ItemID      int     `json:"item_id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Cost        int     `json:"cost"`
	FlingPower  *int    `json:"fling_power"`
	FlingEffect *string `json:"fling_effect"`
}
