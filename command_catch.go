package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {

	if len(args) == 0 {
		return errors.New("missing pokemon name")
	}
	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	rnd := rand.Intn(pokemon.BaseExperience)
	if rnd > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Printf("You may now use the 'pokedex' command or 'inspect %s' \nto get more information about this Pokemon\n", pokemon.Name)

	cfg.caughtPokemon[pokemon.Name] = pokemon

	return nil
}
