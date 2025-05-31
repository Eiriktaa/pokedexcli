package models

import "fmt"

type LocationArea struct {
	Name string
	Url  string
}

type LocationListResult struct {
	Count    int
	Next     string
	Previous string
	Results  []LocationArea
}

func PrintLocationData(locations []LocationArea) {
	for _, location := range locations {
		fmt.Println(location.Name)
	}
}
