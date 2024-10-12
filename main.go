package main

import (
	"pokedex/cli"
	"pokedex/internal/api"
	"time"
)

func main() {
	config := cli.Config{
		Client:         api.NewClient(5 * time.Minute),
		CaughtPokemons: make(map[string]api.PokemonResp),
	}

	cli.StartREPL(&config)
}
