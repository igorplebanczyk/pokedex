package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type commands struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]commands {
	return map[string]commands{
		"map": {
			name:        "map",
			description: "Shows a list of next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows a list of previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Shows a list of pokemons in a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a pokemon",
			callback:    commandCatch,
		},
		"help": {
			name:        "help",
			description: "Shows a list of commands and their usage",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the program",
			callback:    commandExit,
		},
	}
}

func standardizeInput(input string) []string {
	input = strings.ToLower(input)
	return strings.Fields(input)
}

func startREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")

		scanner.Scan()
		userInput := scanner.Text()

		standardizedUserInput := standardizeInput(userInput)
		if len(standardizedUserInput) == 0 {
			continue
		}

		arg := ""
		if len(standardizedUserInput) > 1 {
			arg = standardizedUserInput[1]
		}

		commands := getCommands()
		command, ok := commands[standardizedUserInput[0]]
		if !ok {
			println("Command not found")
			continue
		}

		err := command.callback(cfg, arg)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
	}
}
