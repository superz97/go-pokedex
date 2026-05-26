package main

import (
	"fmt"
	"math/rand"
)

const catchThreshold = 50

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a pokemon name")
	}
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	roll := rand.Intn(pokemon.BaseExperience)
	if roll < catchThreshold {
		fmt.Printf("%s was caught!\n", name)
		cfg.caughtPokemon[name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}
