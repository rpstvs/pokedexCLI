package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(pokemonName string) (bool, error) {

	url := baseURL + "/pokemon/" + pokemonName

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return false, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)

	if err != nil {
		return false, err
	}

	pokemon := Pokemon{}

	err = json.Unmarshal(dat, &pokemon)

	if err != nil {
		return false, err
	}

	return true, nil
}
