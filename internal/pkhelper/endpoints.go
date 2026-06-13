package pkhelper

import (
	"cynthia/pokemon"
	"fmt"
)

func (c *PokemonClient) EndpointGetPokemon(id int) string {
	return fmt.Sprintf("/pokemon/%d", id)
}

func (c *PokemonClient) EndpointGetPokemonSprite(id int, spriteType pokemon.PokemonSprite) string {
	return fmt.Sprintf("/sprites/pokemon/%d/%s", id, spriteType)
}
