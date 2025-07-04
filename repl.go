package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lackingworth/Pokedex-Go/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocation  *string
	prevLocation  *string
	caughtPokemon map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"catch": {
			name:        "catch <pokemon name>",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon name>",
			description: "View details about a caught Pokemon",
			callback:    commandInspect,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"map": {
			name:        "map",
			description: "Get the next page of location areas",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of location areas",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
	}
}

func startRepl(config *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command:", commandName)
			continue
		}

	}
}
