package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Srikumar6529/pokedexcli/pokapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	supportedCommands map[string]cliCommand
	prev              string
	next              string
	client            *pokapi.Client
}

func cleanInput(text string) []string {
	//trimming trailing and leading white space
	trimmedText := strings.TrimSpace(text)

	//converting all the text to lowercase
	loweredText := strings.ToLower(trimmedText)

	//breaking text into words
	words := strings.Fields(loweredText)
	return words
}
func getCommands() map[string]cliCommand {
	supportedCommands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "help for pokedex",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "prints the next 20 locations on the map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "prints the previous 20 locations on the map",
			callback:    commandMapb,
		},
	}
	return supportedCommands
}
func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		err := scanner.Scan()
		if err != false {
			scanner.Err()
		}
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}
		commandName := cleanInput(text)[0]
		//fmt.Println(command)

		command, exists := cfg.supportedCommands[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			fmt.Println("Unsupported Command")
			continue
		}

	}
}
