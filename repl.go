package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/superz97/go-pokedex/internal/pokeapi"
	"github.com/superz97/go-pokedex/internal/pokecache"
)

type config struct {
	Next          *string
	Previous      *string
	pokeapiClient pokeapi.Client
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	cache := pokecache.NewCache(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(cache),
	}
	commands := getCommands()
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		if cmd, ok := commands[commandName]; ok {
			if err := cmd.callback(cfg, words[1:]); err != nil {
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
		"explore": {
			name:        "explore",
			description: "Explore a location area for Pokemon",
			callback:    commandExplore,
		},
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
