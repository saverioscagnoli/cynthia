package pokemon

import "fmt"

func (c *PokemonClient) EndpointGetPokemon(id int) string {
	return fmt.Sprintf("%s/pokemon/%d", c.baseURL, id)
}
