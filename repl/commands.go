package repl

import (
	"fmt"
	"os"
)

func helpCommand (config *Config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to Pokedex!")
	fmt.Println()
	fmt.Println("Usage: ")
	
	for _, cmd := range getCommands() {
		fmt.Println(cmd.name, ":",  cmd.description)
	}
	fmt.Println()
	return nil
}

func exitCommand(config *Config, args ...string) error {
	os.Exit(0)
	return nil
}


func mapCommand(config *Config, args ...string) error {
	locations, err := config.PokeJourney.GoNext()
	
	if err != nil {
		return err
	}

	for _, location := range locations {
		fmt.Println(location)
	}

	return nil
}

func mapbCommand(config *Config, args ...string) error {
	locations, err := config.PokeJourney.GoBack()
	if err != nil {
		return err
	}

	for _, location := range locations {
		fmt.Println(location)
	}

	return nil
}

func exploreCommand(config *Config, args ...string) error {
	area := args[0]
	fmt.Println("Exploring", area, "...")
	pokemons, err := config.PokeJourney.LookForPokemonsInArea(args[0])
	
	if err != nil {
		return err
	}

	fmt.Println("Found pokemons:")

	for _, pokemon := range pokemons {
		fmt.Println(" - ", pokemon)
	}

	return nil
}

func catchCommand(config *Config, args ...string) error {
	pokemonName := args[0]
	
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)

	result, err := config.PokeJourney.TryCatch(pokemonName)

	fmt.Println(result)

	if err != nil {
		return err
	}

	return nil
}

func inspectCommand(config *Config, args ...string) error {
	name := args[0]
	config.PokeJourney.InspectOwnedPokemon(name)

	return nil
}

func pokedexCommand(config *Config, args ...string) error {
	pokemons := config.PokeJourney.ListOwnedPokemonNames()

	for _, pokemon := range pokemons {
		fmt.Println(" -", pokemon)
	}

	return nil
}


type cliCommand struct {
	name        string
	description string
	callback    func(config *Config, args ...string) error
}

func getCommands () map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: helpCommand,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: exitCommand,
		},
		"map": {
			name: "map",
			description: "Explore next 20 locations in pokemon world",
			callback: mapCommand,
		},
		"mapb": {
			name: "mapb",
			description: "Explore previous 20 locations in pokemon world",
			callback: mapbCommand,
		},
		"explore": {
			name: "explore <location_name>",
			description: "Search for pokemons in the area",
			callback: exploreCommand,
		},
		"catch": {
			name: "catch <pokemon_name>",
			description: "Try to catch a pokemon",
			callback: catchCommand,
		},
		"inspect": {
			name: "inspect <pokemon_name>",
			description: "View details about a caught Pokemon",
			callback: inspectCommand,
		},
		"pokedex": {
			name: "pokedex",
			description: "See all the pokemon you've caught",
			callback: pokedexCommand,
		},
	}
}