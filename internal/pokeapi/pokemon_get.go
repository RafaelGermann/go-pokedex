package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if val, exists := c.cache.Get(url); exists {
		locationsResp := Pokemon{}
		if err := json.Unmarshal(val, &locationsResp); err != nil {
			return Pokemon{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	if err = json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)
	return pokemon, nil
}
