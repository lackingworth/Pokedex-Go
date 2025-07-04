package main

import (
	"errors"
	"fmt"
)

func commandMapf(config *config) error {
	locationsRes, err := config.pokeapiClient.ListLocations(config.nextLocation)
	if err != nil {
		return err
	}

	config.nextLocation = locationsRes.Next
	config.prevLocation = locationsRes.Previous

	for _, loc := range locationsRes.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(config *config) error {
	if config.prevLocation == nil {
		return errors.New("you're on the first page")
	}

	locationsRes, err := config.pokeapiClient.ListLocations(config.prevLocation)
	if err != nil {
		return err
	}

	config.nextLocation = locationsRes.Next
	config.prevLocation = locationsRes.Previous

	for _, loc := range locationsRes.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
