package internals

import (
	"fmt"
	"pokedexcli/internals/configs/context"
	"pokedexcli/internals/models"
)

func GetMapDataForward(r *context.ReplContext) error {
	var location models.LocationListResult
	if err := getMapdata(r, r.Next+locEndpoint, &location); err != nil {
		return err
	}
	handleMapResult(r, location)
	return nil
}
func GetMapDataBackward(r *context.ReplContext) error {

	var location models.LocationListResult
	if err := getMapdata(r, r.Previous+locEndpoint, &location); err != nil {
		return err
	}
	handleMapResult(r, location)
	return nil
}
func handleMapResult(r *context.ReplContext, res models.LocationListResult) {

	models.PrintLocationData(res.Results)

	r.Next = res.Next
	r.Previous = res.Previous
}

func GetExploreData(r *context.ReplContext, city string) error {

	var res models.ExploreResponse
	err := getMapdata(r, r.BaseUrl+locEndpoint+city, &res)
	if err != nil {
		return err
	}
	res.LogPossiblePokemonEncounter()
	return nil
}

// to be called by getMapForward,getMapBackward, getExploredata
func getMapdata[T any](r *context.ReplContext, url string, data *T) error {

	if url == "" {
		fmt.Println("url, missing failed to read new data")
		return nil
	}
	err := HttpGetApiDataWithUnmarshal(r, url, data)
	if err != nil {
		fmt.Errorf("something went wrong with getting location data %v", err)
		return err
	}

	return nil
}
