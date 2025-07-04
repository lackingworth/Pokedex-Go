package main

import (
	"fmt"
)

func commandHelp(config *config, args ...string) error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
