package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(url *string) (PokemonResp, error) {
	fullUrl := baseURL + "/pokemon/" + *url

	data, ok := c.cache.Get(fullUrl)
	if ok {
		pokemonResp := PokemonResp{}
		err := json.Unmarshal(data, &pokemonResp)
		if err != nil {
			return PokemonResp{}, err
		}

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return PokemonResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return PokemonResp{}, fmt.Errorf("no such pokemon exists")
	}

	if resp.StatusCode > 299 {
		return PokemonResp{}, fmt.Errorf("received status code %d", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResp{}, err
	}

	pokemonResp := PokemonResp{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return PokemonResp{}, err
	}

	c.cache.Add(fullUrl, data)

	return pokemonResp, nil
}
