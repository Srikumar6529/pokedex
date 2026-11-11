package main

import (
	"fmt"
	"time"

	"github.com/Srikumar6529/pokedexcli/pokapi"
)

func main() {
	fmt.Println("Welcome to PokedexCLI")
	cfg := config{
		supportedCommands: getCommands(),
		prev:              "",
		next:              "https://pokeapi.co/api/v2/location-area/",
		client:            pokapi.NewClient(5 * time.Minute),
	}
	startRepl(&cfg)
}
