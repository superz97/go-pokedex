package main

import "fmt"

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a location area name")
	}
	area := args[0]
	url := "https://pokeapi.co/api/v2/location-area/" + area

	fmt.Printf("Exploring %s...\n", area)
	result, err := cfg.pokeapiClient.GetLocationAreaDetail(url)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range result.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
