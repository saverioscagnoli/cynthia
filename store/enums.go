package store

import "camilla/service/util"

type PokemonColor string

const (
	PokemonColorRed    PokemonColor = "red"
	PokemonColorBlue   PokemonColor = "blue"
	PokemonColorYellow PokemonColor = "yellow"
	PokemonColorGreen  PokemonColor = "green"
	PokemonColorBlack  PokemonColor = "black"
	PokemonColorBrown  PokemonColor = "brown"
	PokemonColorPurple PokemonColor = "purple"
	PokemonColorGray   PokemonColor = "gray"
	PokemonColorWhite  PokemonColor = "white"
	PokemonColorPink   PokemonColor = "pink"
)

func (c *PokemonColor) UnmarshalJSON(b []byte) error {
	return util.UnmarshalString(b, c)
}

type MoveLearnMethod string

const (
	MoveLearnMethodLevelUp MoveLearnMethod = "LevelUp"
	MoveLearnMethodEgg     MoveLearnMethod = "Egg"
	MoveLearnMethodTutor   MoveLearnMethod = "Tutor"
	MoveLearnMethodMachine MoveLearnMethod = "Machine"
)

func (m *MoveLearnMethod) UnmarshalJSON(b []byte) error {
	return util.UnmarshalString(b, m)
}

type MoveDamageClass string

const (
	MoveDamageClassPhysical MoveDamageClass = "Physical"
	MoveDamageClassSpecial  MoveDamageClass = "Special"
)

func (m *MoveDamageClass) UnmarshalJSON(b []byte) error {
	return util.UnmarshalString(b, m)
}
