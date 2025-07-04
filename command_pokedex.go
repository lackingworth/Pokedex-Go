package main

import (
	"fmt"
)

func commandPokedex(config *config, args ...string) error {
	if config.caughtPokemon == nil {
		return fmt.Errorf("you haven't caught any pokemon yet")
	}
	fmt.Println("Your Pokedex:")
	for pokemonName, _ := range config.caughtPokemon {
		fmt.Printf(" - %s\n", pokemonName)
	}
	return nil
}
