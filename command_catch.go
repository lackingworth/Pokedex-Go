package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon")
	}

	name := args[0]

	_, ok := config.caughtPokemon[name]
	if ok {
		return errors.New("you have already caught that pokemon")
	}

	pokemon, err := config.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)

	config.caughtPokemon[pokemon.Name] = pokemon
	return nil
}
