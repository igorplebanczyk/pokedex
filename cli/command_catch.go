package cli

import (
	"fmt"
	"math"
	"math/rand"
)

func catchProbability(baseXP int) float64 {
	const k = -0.01
	const midpoint = 162 // I don't know anything about Pokémon, but the highest base xp I found was 324

	return 1 / (1 + math.Exp(-k*(float64(baseXP)-midpoint)))
}

func attemptToCatch(baseXP int) bool {
	probability := catchProbability(baseXP)

	randomRoll := rand.Float64()

	return randomRoll <= probability
}

func commandCatch(cfg *Config, pokemon string) error {
	if pokemon == "" {
		return fmt.Errorf("a pokemon name is required")
	}

	resp, err := cfg.Client.GetPokemon(&pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("-------------------------\n")
	fmt.Printf("Base experience: %v\n", resp.BaseExperience)
	fmt.Printf("Catch probability: %.2f\n", catchProbability(resp.BaseExperience))
	fmt.Printf("\nThrowing a pokeball at %v...\n", resp.Name)

	if attemptToCatch(resp.BaseExperience) {
		fmt.Printf("Success! %v was caught and added to your Pokedex!\n", resp.Name)
		cfg.CaughtPokemons[resp.Name] = resp
	} else {
		fmt.Printf("Oh no! %v broke free!\n", resp.Name)
	}

	fmt.Printf("-------------------------\n")

	return nil
}
