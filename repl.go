package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jocdiazmu/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    *pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{

		"catch": {
			name:        "catch <pokemon name>",
			description: "Try to catch a Pokemon with a given name",
			callback:    commandCatch,
		},

		"inspect": {
			name:        "inspect <pokemon name>",
			description: "Get the stats of a Pokemen with a given name",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays your pokedex",
			callback:    commandPokedex,
		},
		"explore": {
			name:        "explore <location>",
			description: "Explores a location",
			callback:    commandExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exit the CLI",
			callback:    commandExit,
		},

		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 Pokemon locations map",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 Pokemon locations map",
			callback:    commandMapb,
		},
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := commands[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
