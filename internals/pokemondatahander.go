package internals

import (
	"fmt"
	"math/rand"
	"pokedexcli/internals/configs/context"
	"pokedexcli/internals/models"
)

func getPokemonByName(r *context.ReplContext, pokemonName string) (models.Pokemon, error) {

	var res models.Pokemon
	err := HttpGetApiDataWithUnmarshal(r, r.BaseUrl+pokemonEndpoint+pokemonName, &res)
	if err != nil {
		return res, err
	}
	// res.LogPossiblePokemonEncounter()
	return res, nil
}

func checkIfSuccsessfulCatch(pokedex models.Pokedex, pokemon models.Pokemon) {
	const escapeThreshold = 500
	escapePower := rand.Intn(escapeThreshold) + pokemon.BaseExperience
	if escapePower > escapeThreshold {
		fmt.Println("Ohh no, the pokemon escaped")
		return
	}
	pokedex.Caught[pokemon.Name] = pokemon
	fmt.Println("caught ", pokemon.Name)
}

func CatchPokemon(r *context.ReplContext, pokemonName string) error {
	fmt.Println(fmt.Sprintf("Throwing a Pokeball at %s...", pokemonName))
	pokemon, ok := r.Pokedex.Seen[pokemonName]
	var err error
	if !ok {
		pokemon, err = getPokemonByName(r, pokemonName)
		if err != nil {
			return err
		}
	}
	r.Pokedex.Seen[pokemonName] = pokemon
	checkIfSuccsessfulCatch(r.Pokedex, pokemon)
	return nil
}

func InspectPokemon(r *context.ReplContext, pokemonName string) {
	pokemon, ok := r.Pokedex.Seen[pokemonName]
	if !ok {
		fmt.Println(pokemonName, " not seen before. Unable to display stats")
		return
	}

	pokemon.DisplayStats()
}

// // to be called by getMapForward,getMapBackward, getExploredata
// func getMapdata[T any](r *context.ReplContext, url string, data *T) error {
//
// 	if url == "" {
// 		fmt.Println("url, missing failed to read new data")
// 		return nil
// 	}
// 	err := HttpGetLocationAreas(r, url, data)
// 	if err != nil {
// 		fmt.Errorf("something went wrong with getting location data %v", err)
// 		return err
// 	}
//
// 	return nil
// }
