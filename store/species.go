package store

type PokemonSpecies struct {
	ID                int          `json:"id"`
	Name              string       `json:"name"`
	NationalDexNumber int          `json:"national_dex_number"`
	GenderChance      int          `json:"gender_chance"`
	CaptureRate       int          `json:"capture_rate"`
	BaseHappiness     int          `json:"base_happiness"`
	IsBaby            bool         `json:"is_baby"`
	IsLegendary       bool         `json:"is_legendary"`
	IsMythic          bool         `json:"is_mythic"`
	HatchCounter      int          `json:"hatch_counter"`
	GenderDifferences bool         `json:"gender_differences"`
	FormsSwitchable   bool         `json:"forms_switchable"`
	Color             PokemonColor `json:"color"`
}

type Type struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
