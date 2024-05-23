package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("expected a pokemon name")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.CatchPokemon(pokemonName)

	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Trying to catch %s ...\n", pokemon.Name)
	if res > 40 {
		fmt.Printf("Congratulations you've catched a Pokemon\n")
		cfg.Pokedex[pokemon.Name] = pokemon

	} else {
		fmt.Printf("%s got away!! Try again\n", pokemon.Name)
	}

	return nil
}
