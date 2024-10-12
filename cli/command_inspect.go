package cli

import (
	"fmt"
)

func commandInspect(cfg *Config, name string) error {
	if name == "" {
		return fmt.Errorf("no pokemon name provided")
	}

	pokemon, ok := cfg.CaughtPokemons[name]
	if !ok {
		return fmt.Errorf("%v was not found in caught pokemons", name)
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Stats: ")
	for _, stat := range pokemon.Stats {
		fmt.Printf("%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types: ")
	for _, t := range pokemon.Types {
		fmt.Printf("%v\n", t.Type.Name)
	}

	return nil
}
