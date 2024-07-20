package main

import "fmt"

func commandExplore(cfg *config, location string) error {
	if location == "" {
		return fmt.Errorf("location is required")
	}

	endpoint := "/location-area/" + location
	resp, err := cfg.pokeapiClient.GetLocationAreaDetails(&endpoint)
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
