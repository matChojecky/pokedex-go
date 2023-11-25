package pokeJourney

import "github.com/matChojecky/pokedexcli/pokeapi"


type PokeJourney struct {
	pokeApi *pokeapi.PokeApi
	nextArea string
	prevArea string
	owned map[string]Pokemon
}


func BeginJourney(pokeApi *pokeapi.PokeApi) PokeJourney {
	return PokeJourney {
		pokeApi: pokeApi,
		nextArea: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		prevArea: "",
		owned: make(map[string]Pokemon),
	}
}