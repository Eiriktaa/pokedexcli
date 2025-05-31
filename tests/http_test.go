package main

import (
	"fmt"
	"pokedexcli/internals"
	"pokedexcli/internals/configs/context"
	"pokedexcli/internals/models"
	"pokedexcli/internals/pokecache"
	"testing"
	"time"
)

func TestHttpGetLocation(t *testing.T) {
	baseurl := "https://pokeapi.co/api/v2/"
	locationEndpoint := "location-area/"

	ctx := context.NewReplContext(baseurl + locationEndpoint)
	var list models.LocationListResult
	err := internals.HttpGetLocationAreas(&ctx, baseurl+locationEndpoint, &list)
	if err != nil {
		t.Errorf("get request failed: %v ", err)
	}
	if len(list.Results) != 20 {
		t.Errorf("expected 20 elements in standard reqeuest, check values: %d ", len(list.Results))
	}

	firstCityName := "canalave-city-area"
	if list.Results[0].Name != firstCityName {
		t.Errorf("found wrong city in the first index: expected %s found %s", firstCityName, list.Results[0].Name)
	}

}

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := pokecache.NewCache(interval)
			cache.AddData(c.key, c.val)
			val, ok := cache.GetData(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := pokecache.NewCache(baseTime)
	cache.AddData("https://example.com", []byte("testdata"))

	_, ok := cache.GetData("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.GetData("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
