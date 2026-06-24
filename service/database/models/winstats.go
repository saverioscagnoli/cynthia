package models

type WinStats struct {
	Wins    int     `json:"wins"`
	Losses  int     `json:"losses"`
	Draws   int     `json:"draws"`
	Winrate float32 `json:"winrate"`

	CurrentStreak int `json:"current_streak"`
	BestStreak    int `json:"best_streak"`

	SingleWins int `json:"single_wins"`
	DoubleWins int `json:"double_wins"`

	TotalDamageDealt    int `json:"total_damage_dealt"`
	TotalDamageReceived int `json:"total_damage_received"`
	PokemonFainted      int `json:"pokemon_fainted"`
	PokemonLost         int `json:"pokemon_lost"`
}
