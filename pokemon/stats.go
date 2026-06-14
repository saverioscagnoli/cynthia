package pokemon

type StatName string

type Stat struct {
	ID   int      `json:"id"`
	Name StatName `json:"name"`
}

type PokemonStat struct {
	Stat
	Value int `json:"value"`
}
