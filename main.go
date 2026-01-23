package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for x := 1; x > 0; {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		wordList := cleanInput(input)
		if len(wordList) < 1 {
			continue
		}
		fmt.Printf("Your command was: %s\n", wordList[0])

	}
}
