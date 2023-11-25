package pokeJourney

import (
	"errors"

	"github.com/matChojecky/pokedexcli/pokeapi"
)

func (j *PokeJourney) GoNext() ([]string, error) {
	response, err := j.pokeApi.ListLocationAreas(j.nextArea)
	if err != nil {
		return []string{}, err
	}

	j.setNewDestinations(response.Next, response.Previous)

	return getLocationAreaNamesFromResponse(response), nil
}

func (j *PokeJourney) GoBack() ([]string, error) {
	if j.prevArea == "" {
		return []string{}, errors.New("just and error message")
	}
	response, err := j.pokeApi.ListLocationAreas(j.prevArea)

	if err != nil {
		return []string {}, nil
	}

	j.setNewDestinations(response.Next, response.Previous)

	return getLocationAreaNamesFromResponse(response), nil
}

func (j *PokeJourney) setNewDestinations(next , prev string) {
	j.nextArea = next
	j.prevArea = prev
}


func getLocationAreaNamesFromResponse(response pokeapi.PokeLocationsResponse) []string {
	locNames := []string{};

	for _, loc := range response.Results {
		locNames = append(locNames, loc.Name)
	}

	return locNames
}

func (j *PokeJourney) LookForPokemonsInArea(area string) ([]string, error) {
	response, err := j.pokeApi.GetAreaByName(area)
	if err != nil {
		return []string{}, err
	}

	pokemonsInArea := []string {}

	for _, encounter := range response.PokemonEncounters {
		pokemonsInArea = append(pokemonsInArea, encounter.Pokemon.Name)
	}

	return pokemonsInArea, nil
}