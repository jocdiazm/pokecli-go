package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return fmt.Errorf("you haven't caught any Pokemon yet")
	}

	fmt.Println("Your pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
