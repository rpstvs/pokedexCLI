package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	for {
		reader := bufio.NewScanner(os.Stdin)

		fmt.Print("Pokedex > ")

		reader.Scan()

		words := cleanInput(reader.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]

		if exists {
			err := command.callback()

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
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the program",
			callback:    commandExit,
		},
	}
}
