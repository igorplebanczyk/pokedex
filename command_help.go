package main

import "fmt"

func commandHelp() error {
	commands := getCommands()
	for _, v := range commands {
		fmt.Printf("- %v: %v\n", v.name, v.description)
	}
	return nil
}
