package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rpstvs/pokedexCLI/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	for {
		reader := bufio.NewScanner(os.Stdin)

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

		command, exists := getCommands()[commandName]

		if exists {
			err := command.callback(cfg, args...)

			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown Command")
			continue
		}

	}

}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get map locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous map locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a certain area for pokemons",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the program",
			callback:    commandExit,
		},
	}
}
