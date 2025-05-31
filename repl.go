package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internals/configs"
	"pokedexcli/internals/configs/context"
	"pokedexcli/internals/helpers"
)

func startRepl() {
	replContext := context.NewReplContext("https://pokeapi.co/api/v2/")
	scanner := bufio.NewScanner(os.Stdin)
	validcmds := configs.GenerateCMDS()
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		cmds := helpers.CleanInput(scanner.Text())
		if len(cmds) == 0 {
			continue
		}
		inputCMD := cmds[0]
		clicmd, ok := validcmds[inputCMD]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := clicmd.Callback(&replContext, cmds...)
		if err != nil {
			fmt.Println(fmt.Errorf("error found: %v", err))
		}
	}
}
