package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokeapi/pokecache"
	"strings"
)

func cleanInput(text string) []string {
	var stringList []string
	var nextWord string
	for i, character := range text {
		if character == ' ' && len(nextWord) > 0 {
			stringList = append(stringList, strings.ToLower(nextWord))
			nextWord = ""
			continue
		}
		if character == ' ' {
			continue
		}
		nextWord += string(character)
		if i == len(text)-1 {
			stringList = append(stringList, strings.ToLower(nextWord))
		}
	}
	return stringList
}

func startREPL(cache *pokecache.Cache, pkDex *dex) {
	scanner := bufio.NewScanner(os.Stdin)
	var configFile config

	for x := 1; x > 0; {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		wordList := cleanInput(input)
		if len(wordList) < 1 {
			continue
		}
		//fmt.Printf("Your command was: %s\n", wordList[0])
		commandName := wordList[0]
		parameter := ""
		if len(wordList) > 1 {
			parameter = wordList[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cache, &configFile, pkDex, parameter)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Prints 20 locations progressively",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Prints the last 20 locations from map",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Lists pokemon found in a given area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "attempt to catch a pokemon",
			callback:    commandCatch,
		},
		/*"scan": {
			name:        "scan",
			description: "scanning for max base exp value",
			callback:    baseExpScan,
		},*/
	}
}
