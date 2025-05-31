package models

type Pokedex struct {
	Caught map[string]Pokemon
	Seen   map[string]Pokemon
}

func NewPokedex() Pokedex {
	return Pokedex{
		Caught: make(map[string]Pokemon),
		Seen:   make(map[string]Pokemon),
	}
}
