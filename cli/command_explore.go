package cli

import (
	"fmt"
)

func commandExplore(cfg *Config, location string) error {
	if location == "" {
		return fmt.Errorf("location is required")
	}

	resp, err := cfg.Client.GetLocationAreaDetails(&location)
	if err != nil {
		return err
	}

	fmt.Printf("-------------------------")
	fmt.Printf("\nPokemons in location: \n")
	for _, pokemon := range resp.PokemonEncounters {
		println(pokemon.Pokemon.Name)
	}
	fmt.Printf("-------------------------\n")

	return nil
}
