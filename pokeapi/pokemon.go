package pokeapi

import (
	"encoding/json"
	"fmt"
)

type GetPokemonResponse struct {
	BaseExperience int `json:"base_experience"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Height int `json:"height"`
	Weight int `json:"weight"`
}

func (c *PokeApi) GetPokemonByName(name string) (GetPokemonResponse, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", name)
	res := GetPokemonResponse{}
	cached, ok := c.cache.Get(url)
	if ok {
		err := json.Unmarshal(cached, &res)

		return res, err
	}
	body, err := c.httpGetWithCache(url)
	
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(body, &res)

	return res, err

}


