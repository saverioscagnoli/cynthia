package pkapi

type Item struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Cost        int     `json:"cost"`
	FlingPower  *int    `json:"fling_power"`
	FlingEffect *string `json:"fling_effect"`
}

type HeldItem struct {
	Item
	Rarity int `json:"rarity"`
}
