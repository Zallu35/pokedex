package main

import "strings"

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
