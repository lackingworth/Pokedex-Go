package main

import (
	"time"

	"github.com/lackingworth/Pokedex-Go/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	config := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeClient,
	}

	startRepl(config)
}
