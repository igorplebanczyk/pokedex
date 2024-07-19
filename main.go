package main

import "github.com/igorplebanczyk/pokedex/pokeapi"

type config struct {
	pokeapiClient            pokeapi.Client
	nextLocationAreasURL     *string
	previousLocationAreasURL *string
}

func main() {
	config := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	startREPL(&config)
}
