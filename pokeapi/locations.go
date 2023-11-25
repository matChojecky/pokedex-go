package pokeapi

import (
	"encoding/json"
	"fmt"
)

type PokeLocationsResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokeLocationAreaByIdResponse struct {
	Name string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *PokeApi) ListLocationAreas(url string) (PokeLocationsResponse, error) {
	cached, ok := c.cache.Get(url)
	result := PokeLocationsResponse{}
	var data []byte 
	if ok {
		err := json.Unmarshal(cached, &result)

		return result, err
	}

	body, err := c.httpGetWithCache(url)
	if err != nil {
		return result, err
	}

	data = body

	err = json.Unmarshal(data, &result)

	return result, err
}

func (c *PokeApi) GetAreaByName(location string) (PokeLocationAreaByIdResponse, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v", location)
	result := PokeLocationAreaByIdResponse{}
	cached, ok := c.cache.Get(url)
	
	if ok {
		err := json.Unmarshal(cached, &result)

		return result, err
	}

	body, err := c.httpGetWithCache(url)
	
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)

	return result, err
}