package main

import (
	"github.com/igorplebanczyk/pokedex/internal/pokeapi"
	"time"
)

type config struct {
	pokeapiClient            pokeapi.Client
	nextLocationAreasURL     *string
	previousLocationAreasURL *string
	caughtPokemons           map[string]pokeapi.PokemonResp
}

func main() {
	config := config{
		pokeapiClient:  pokeapi.NewClient(5 * time.Minute),
		caughtPokemons: make(map[string]pokeapi.PokemonResp),
	}

	startREPL(&config)
}
