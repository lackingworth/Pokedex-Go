package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Provide usage manual",
		callback:    commandHelp,
	},
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	splitString := strings.Fields(lowerCase)
	return splitString
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
	return nil
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex: > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		command := words[0]

		if len(words) == 0 {
			continue
		}

		switch command {
		case commands["exit"].name:
			err := commandExit()
			if err != nil {
				fmt.Println(err)
			}
		case commands["help"].name:
			err := commandHelp()
			if err != nil {
				fmt.Println(err)
			}
		default:
			fmt.Println("Unknown command: " + command)
		}

	}
}
