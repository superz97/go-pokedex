package pokeapi

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	var result Pokemon
	if err := c.get("https://pokeapi.co/api/v2/pokemon/"+name, &result); err != nil {
		return result, err
	}
	return result, nil
}
