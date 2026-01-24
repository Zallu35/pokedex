package main

import (
	"bufio"
	"fmt"
	"os"
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

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)
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

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
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

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println(`
	Welcome to the Pokedex!
	Usage:
	`)
	commands := getCommands()
	for i := range commands {
		fmt.Printf("%s: %s\n", i, commands[i].description)
	}
	return nil
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
	}
}
