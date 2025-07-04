package main

import (
	"errors"
	"fmt"
)

func commandExplore(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("please provide a location")
	}

	name := args[0]
	location, err := config.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}
