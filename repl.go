package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/superz97/go-pokedex/internal/pokecache"
)

type config struct {
	Next     *string
	Previous *string
	cache    *pokecache.Cache
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	cfg := &config{
		cache: pokecache.NewCache(5 * time.Second),
	}
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		if cmd, ok := commands[commandName]; ok {
			if err := cmd.callback(cfg); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    commandMapb,
		},
	}
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/?limit=20"
	if cfg.Next != nil {
		url = *cfg.Next
	}

	result, err := fetchLocationAreas(url, cfg.cache)
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

func commandMapb(cfg *config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	result, err := fetchLocationAreas(*cfg.Previous, cfg.cache)
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

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
