package main

import (
	"time"

	"github.com/RafaelGermann/pokedex-go/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
