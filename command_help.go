package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for key, val := range cfg.supportedCommands{
		fmt.Println(key,":",val.description)
	}
	return nil
}
