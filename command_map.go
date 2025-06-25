package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = resp.Next
	cfg.prevLocationURL = resp.Previous

	for _, l := range resp.Results {
		fmt.Println(l.Name)
	}
	return nil
}

func commandMapB(cfg *config, args ...string) error {
	if cfg.prevLocationURL == nil {
		return errors.New("you're on the first page")
	}

	resp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = resp.Next
	cfg.prevLocationURL = resp.Previous

	for _, l := range resp.Results {
		fmt.Println(l.Name)
	}
	return nil
}
