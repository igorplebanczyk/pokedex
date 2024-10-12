package cli

import "pokedex/internal/api"

type Config struct {
	Client                   api.Client
	CaughtPokemons           map[string]api.PokemonResp
	nextLocationAreasURL     *string
	previousLocationAreasURL *string
}
