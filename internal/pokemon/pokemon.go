package pokemon

import (
	"cynthia/cmd/pkapi"
	"cynthia/util"
)

func (c *PokemonClient) GetPokemon(id int) (*Pokemon, error) {
	res, err := c.get(c.EndpointGetPokemon(id))

	if err != nil {
		return nil, err
	}

	return util.Decode[Pokemon](res)
}

func (c *PokemonClient) GetPokemonSprite(pkapi.PokemonSprite)
