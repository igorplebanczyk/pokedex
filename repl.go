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
	callback    func() error
}

func getCommands() map[string]commands {
	return map[string]commands{
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

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex> ")

		scanner.Scan()
		userInput := scanner.Text()

		standardizedUserInput := standardizeInput(userInput)
		if len(standardizedUserInput) == 0 {
			continue
		}

		commands := getCommands()
		command, ok := commands[standardizedUserInput[0]]
		if !ok {
			println("Command not found")
			continue
		}

		err := command.callback()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	}
}
