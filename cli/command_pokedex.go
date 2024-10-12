package cli

import (
	"fmt"
)

func commandPokedex(cfg *Config, _ string) error {
	if len(cfg.CaughtPokemons) == 0 {
		fmt.Printf("No pokemons caught yet\n")
		return nil
	}

	fmt.Printf("-------------------------")
	fmt.Printf("\nYour Pokedex: \n")
	for _, pokemon := range cfg.CaughtPokemons {
		fmt.Printf("- %v\n", pokemon.Name)
	}
	fmt.Printf("-------------------------\n")

	return nil
}
