package pokeJourney

import (
	"fmt"
	"math/rand"
)


type Pokemon struct {
	Name string
	Height int
	Weight int
	Stats []struct{
		Base int
		Stat string
	}
	Types []string
}

func (j *PokeJourney) TryCatch(pokemon string) (string, error) {
	pokemonData, err := j.pokeApi.GetPokemonByName(pokemon)
	if err != nil {
		return fmt.Sprintf("%v escaped!", pokemon), err
	}

	chance := rand.Int() * 2
	if pokemonData.BaseExperience > chance {
		return fmt.Sprintf("%v escaped!", pokemon), nil	
	}

	types := []string{}

	for _, typeSlot := range pokemonData.Types {
		types = append(types, typeSlot.Type.Name)
	}

	stats := []struct {
		Base int
		Stat string
}{}

for _, stat := range pokemonData.Stats {
	stats = append(stats, struct{Base int; Stat string}{
		Base: stat.BaseStat,
		Stat: stat.Stat.Name,
	})
}

	j.owned[pokemon] = Pokemon{
		Name: pokemonData.Name,
		Height: pokemonData.Height,
		Weight: pokemonData.Weight,
		Types: types,
		Stats: stats,
	}

	fmt.Println("You may now inspect it with the inspect command.")

	return fmt.Sprintf("%v was caught!", pokemon), err
}

func (j *PokeJourney) ListOwnedPokemonNames() []string {
	names := []string{}

	for _, pokemon := range j.owned {
		names = append(names, pokemon.Name)
	}

	return names
}

func (j *PokeJourney) InspectOwnedPokemon(name string) {
	pokemon, ok := j.owned[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats: ")
	
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat, stat.Base)
	}

	fmt.Println("Types: ")
	for _, pokeType := range pokemon.Types {
		fmt.Println("  -", pokeType)
	}

}