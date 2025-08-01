package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	resp, err := cfg.pokeapiClient.GetLotication(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", resp.Name)
	fmt.Println("Found Pokemon:")
	for _, l := range resp.PokemonEncounters {
		fmt.Println("- " + l.Pokemon.Name)
	}
	return nil
}
