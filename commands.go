package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"pokedex/internal/pokeapi/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pokecache.Cache, *config, *dex, string) error
}

type config struct {
	Next     string
	Previous string
}

func commandExit(cache *pokecache.Cache, cf *config, a *dex, unused string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cache *pokecache.Cache, cf *config, a *dex, unused string) error {
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

func fetchData(cache *pokecache.Cache, url string) ([]byte, error) {
	data, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("http response error: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return nil, fmt.Errorf("response failed with status code %d and \nbody: %s\n", res.StatusCode, data)
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("io read error %v", err)
		}

		cache.Add(url, data)
	}
	return data, nil
}

/*
// I made this function purely to identify the highest base exp value among all pokemon.
// Highest base exp value is blissy at 608 base exp.

func baseExpScan(cache *pokecache.Cache, cf *config, unused string) error {
	highestValue := 0
	pkmnName := ""
	for i := 1; i < 1026; i++ {
		fmt.Printf("%d/1025", i) //printing a progress meter
		url := "https://pokeapi.co/api/v2/pokemon/" + fmt.Sprintf("%d", i)
		res, err := http.Get(url)
		if err != nil {
			fmt.Printf("http response error: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			fmt.Printf("response failed with status code %d and \nbody: %s\n", res.StatusCode, res.Body)
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("io read error %v", err)
		}

		pokeData := struct {
			BaseExperience int    `json:"base_experience"`
			Name           string `json:"name"`
		}{}
		if err := json.Unmarshal(data, &pokeData); err != nil {
			fmt.Printf("error unmarshaling json: %v", err)
		}
		if pokeData.BaseExperience > highestValue {
			highestValue = pokeData.BaseExperience
			pkmnName = pokeData.Name
		}
		fmt.Print("\033[H\033[2J") //this was an attempt to make the progress meter not flood the CLI. Unfortunately this clears the entire terminal, not just the line it's on.
	}
	fmt.Printf("Name: %s\nBase EXP: %v\n", pkmnName, highestValue)
	return nil
}*/
