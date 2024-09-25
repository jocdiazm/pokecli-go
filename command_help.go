package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex CLI!")
	fmt.Println("\nAvailable commands:")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
