package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("expected pokemon name")
	}

	name := args[0]

	pokemon, ok := cfg.Pokedex[name]

	if !ok {

		fmt.Printf("you haven't caught %s", name)

	}
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println("-", typeInfo.Type.Name)
	}

	return nil
}
