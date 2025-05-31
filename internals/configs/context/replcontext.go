package context

import (
	"pokedexcli/internals/models"
	"pokedexcli/internals/pokecache"
	"time"
)

type ReplContext struct {
	Next     string
	BaseUrl  string
	Previous string
	Cache    pokecache.PokeCache
	Pokedex  models.Pokedex
}

func NewReplContext(initalUrl string) ReplContext {
	return ReplContext{
		Next:     initalUrl,
		BaseUrl:  initalUrl,
		Previous: "",
		Cache:    pokecache.NewCache(time.Second * 5),
		Pokedex:  models.NewPokedex(),
	}
}
