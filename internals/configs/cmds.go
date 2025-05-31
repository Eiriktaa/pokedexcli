package configs

import (
	"fmt"
	"os"
	"pokedexcli/internals"
	"pokedexcli/internals/configs/context"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*context.ReplContext, ...string) error
}

func printHelp(cmds map[string]cliCommand) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for key, val := range cmds {
		fmt.Println(key + ": " + val.Description)
	}
	return nil
}
func GenerateCMDS() map[string]cliCommand {
	var cmds map[string]cliCommand
	cmds = map[string]cliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback: func(r *context.ReplContext, arg ...string) error {
				fmt.Println("Closing the Pokedex... Goodbye!")
				os.Exit(0)
				return nil
			},
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    func(r *context.ReplContext, arg ...string) error { return printHelp(cmds) },
		},
		"map": {
			Name:        "map",
			Description: "Gets the 20 next locations",
			Callback: func(r *context.ReplContext, s ...string) error {
				return internals.GetMapDataForward(r)
			},
		},
		"mapb": {
			Name:        "mapb",
			Description: "Gets the 20 previous locations",
			Callback: func(r *context.ReplContext, s ...string) error {

				return internals.GetMapDataBackward(r)
			},
		},
		"pokedex": {
			Name:        "mapb",
			Description: "Gets the 20 previous locations",
			Callback: func(r *context.ReplContext, s ...string) error {
				fmt.Println("Your Pokedex: ")
				for _, pokemon := range r.Pokedex.Caught {
					fmt.Println("- ", pokemon.Name)
				}
				return nil
			}},

		"explore": {
			Name:        "explore",
			Description: "explore <cityname> lists the pokemons you might encounter in the region",
			Callback: func(r *context.ReplContext, args ...string) error {

				if len(args) < 2 {
					fmt.Println("Usage explore <CITYNAME>")
					return nil
				}
				return internals.GetExploreData(r, args[1])
			},
		},
		"catch": {
			Name:        "catch",
			Description: "Attempts to catch the targeted pokemon",
			Callback: func(r *context.ReplContext, args ...string) error {

				if len(args) < 2 {
					fmt.Println("Usage catch <pokemon>")
					return nil
				}
				return internals.CatchPokemon(r, args[1])

			},
		},
		"inspect": {
			Name:        "inspect",
			Description: "Display the targeted pokemon",
			Callback: func(r *context.ReplContext, args ...string) error {

				if len(args) < 2 {
					fmt.Println("Usage inspect <pokemon>")
					return nil
				}
				internals.InspectPokemon(r, args[1])
				return nil

			},
		},
	}
	return cmds
}
