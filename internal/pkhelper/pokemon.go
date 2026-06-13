package pkhelper

import (
	"cynthia/pokemon"
	"cynthia/util"
	"fmt"
	"io"
)

func GetPokemon(id int) (*pokemon.Pokemon, error) {
	res, err := client.get(client.EndpointGetPokemon(id))

	if err != nil {
		return nil, err
	}

	return util.Decode[pokemon.Pokemon](res)
}

func GetPokemonSprite(id int, spriteType pokemon.PokemonSprite) (*[]byte, error) {
	res, err := client.get(client.EndpointGetPokemonSprite(id, spriteType))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("Failed to read sprite: %w", err)
	}

	return &data, nil
}
