package main

import "fmt"

func commandMap(cfg *config, args []string) error {
	url := "https://pokeapi.co/api/v2/location-area/?limit=20"
	if cfg.Next != nil {
		url = *cfg.Next
	}

	result, err := cfg.pokeapiClient.GetLocationAreas(url)
	if err != nil {
		return err
	}

	cfg.Next = result.Next
	cfg.Previous = result.Previous

	for _, area := range result.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapb(cfg *config, args []string) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	result, err := cfg.pokeapiClient.GetLocationAreas(*cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = result.Next
	cfg.Previous = result.Previous

	for _, area := range result.Results {
		fmt.Println(area.Name)
	}
	return nil
}
