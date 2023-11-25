package pokeapi

import (
	"io"
	"net/http"

	"github.com/matChojecky/pokedexcli/cache"
)



type PokeApi struct {
	cache cache.Cache
}

func NewPokeApiClient(pokeCache cache.Cache) PokeApi {
	return PokeApi {
		cache: pokeCache,
	}
}


func(c *PokeApi) httpGetWithCache(url string) ([]byte, error) {
	res, err := http.Get(url)

	if err != nil {
		return []byte{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	
	if err != nil {
		return []byte{}, err
	}

	c.cache.Add(url, body)

	return body, nil
}