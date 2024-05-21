package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(pokemonName string) (Pokemon, error) {

	url := baseURL + "/pokemon/" + pokemonName

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)

	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}

	err = json.Unmarshal(dat, &pokemon)

	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
