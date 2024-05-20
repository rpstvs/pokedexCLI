package main

import (
	"errors"
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("expected a pokemon name")
	}

	pokemonName := args[0]

	catched, err := cfg.pokeapiClient.CatchPokemon(pokemonName)

	if err != nil {
		return err
	}

	fmt.Printf("Trying to catch %s ...\n", pokemonName)
	if catched {
		fmt.Printf("Congratulations you've catched a Pokemon\n")
	} else {
		fmt.Printf("%s got away!! Try again", pokemonName)
	}

	return nil
}
