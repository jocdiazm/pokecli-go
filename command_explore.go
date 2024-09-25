package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("missing location name")
	}
	name := args[0]
	location, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring location %s\n", location.Name)
	fmt.Printf("...Found (%v) Pokemons:\n", len(location.PokemonEncounters))

	for _, pokemon := range location.PokemonEncounters {
		fmt.Printf(" -%s\n", pokemon.Pokemon.Name)
	}

	return nil
}
