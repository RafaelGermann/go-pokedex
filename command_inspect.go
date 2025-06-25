package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, exists := cfg.caughtPokemon[name]
	if !exists {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)

	fmt.Println("Stats:")
	for _, v := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", v.Stat.Name, v.BaseStat)
	}

	fmt.Println("Types:")
	for _, v := range pokemon.Types {
		fmt.Printf("  - %s\n", v.Type.Name)
	}
	return nil
}
