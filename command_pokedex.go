package main

import "fmt"

func commandPokedex(cfg *config, input string) error {
	if len(cfg.caughtPokemons) == 0 {
		fmt.Printf("No pokemons caught yet\n")
		return nil
	}

	fmt.Printf("-------------------------")
	fmt.Printf("\nYour Pokedex: \n")
	for _, pokemon := range cfg.caughtPokemons {
		fmt.Printf("- %v\n", pokemon.Name)
	}
	fmt.Printf("-------------------------\n")

	return nil
}
