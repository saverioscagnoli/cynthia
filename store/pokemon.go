package store

type PokemonSprite string

const (
	PokemonSpriteFront            PokemonSprite = "front"
	PokemonSpriteBack             PokemonSprite = "back"
	PokemonSpriteFrontShiny       PokemonSprite = "front-shiny"
	PokemonSpriteBackShiny        PokemonSprite = "back-shiny"
	PokemonSpriteFrontFemale      PokemonSprite = "front-female"
	PokemonSpriteBackFemale       PokemonSprite = "back-female"
	PokemonSpriteFrontShinyFemale PokemonSprite = "front-shiny-female"
	PokemonSpriteBackShinyFemale  PokemonSprite = "back-shiny-female"
)

var spriteMap = map[string]PokemonSprite{
	"front":              PokemonSpriteFront,
	"back":               PokemonSpriteBack,
	"front-shiny":        PokemonSpriteFrontShiny,
	"back-shiny":         PokemonSpriteBackShiny,
	"front-female":       PokemonSpriteFrontFemale,
	"back-female":        PokemonSpriteBackFemale,
	"front-shiny-female": PokemonSpriteFrontShinyFemale,
	"back-shiny-female":  PokemonSpriteBackShinyFemale,
}

func ParseSprite(s string) (PokemonSprite, bool) {
	v, ok := spriteMap[s]
	return v, ok
}

type PokemonSprites struct {
	Front            *[]byte `json:"front"`
	FrontShiny       *[]byte `json:"front_shiny"`
	Back             *[]byte `json:"back"`
	BackShiny        *[]byte `json:"back_shiny"`
	FrontFemale      *[]byte `json:"front_female"`
	FrontShinyFemale *[]byte `json:"front_shiny_female"`
	BackFemale       *[]byte `json:"back_female"`
	BackShinyFemale  *[]byte `json:"back_shiny_female"`
}

func (s *PokemonSprites) Get(spriteType PokemonSprite) *[]byte {
	switch spriteType {
	case PokemonSpriteFront:
		return s.Front
	case PokemonSpriteFrontShiny:
		return s.FrontShiny
	case PokemonSpriteBack:
		return s.Back
	case PokemonSpriteBackShiny:
		return s.BackShiny
	case PokemonSpriteFrontFemale:
		return s.FrontFemale
	case PokemonSpriteFrontShinyFemale:
		return s.FrontShinyFemale
	case PokemonSpriteBackFemale:
		return s.BackFemale
	case PokemonSpriteBackShinyFemale:
		return s.BackShinyFemale
	default:
		return nil
	}
}

type Pokemon struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	Types       []Type         `json:"types"`
	BaseExp     int            `json:"base_exp"`
	Height      int            `json:"height"`
	DefaultForm bool           `json:"default_form"`
	Weight      int            `json:"weight"`
	Species     PokemonSpecies `json:"species"`
	Moves       []*PokemonMove `json:"moves"`
	Items       []*HeldItem    `json:"items"`
	BaseStats   []*PokemonStat `json:"base_stats"`
}

type PokemonMove struct {
	Move
	LearnMethod MoveLearnMethod `json:"learn_method"`
}

type OwnedPokemon struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Level     int            `json:"level"`
	Types     []Type         `json:"types"`
	Exp       int            `json:"exp"`
	Height    int            `json:"height"`
	Weight    int            `json:"weight"`
	SpeciesID int            `json:"species_id"`
	Moves     []*Move        `json:"moves"` // Max 4
	HeldItem  *Item          `json:"held_item"`
	Stats     []*PokemonStat `json:"stats"`
}
