package main

import (
	"time"

	"github.com/matChojecky/pokedexcli/cache"
	pokeJourney "github.com/matChojecky/pokedexcli/journey"
	"github.com/matChojecky/pokedexcli/pokeapi"
	"github.com/matChojecky/pokedexcli/repl"
)


func main() {
	pokeCache := cache.NewCache(time.Duration(time.Minute * 2))
	pokeApiClient := pokeapi.NewPokeApiClient(&pokeCache);
	pokemonJourney := pokeJourney.BeginJourney(&pokeApiClient)

	config := repl.Config{
		PokeJourney: &pokemonJourney,
	}

	repl.Start(&config)
}
