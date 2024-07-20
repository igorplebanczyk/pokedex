package main

import "fmt"

func commandHelp(cfg *config, arg string) error {
	commands := getCommands()

	fmt.Printf("-------------------------")
	fmt.Printf("\nAvailable commands: \n")
	for _, v := range commands {
		fmt.Printf("- %v: %v\n", v.name, v.description)
	}
	fmt.Printf("-------------------------\n")

	return nil
}
