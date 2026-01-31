package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"pokedex/internal/pokeapi/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pokecache.Cache, *config, string) error
}

type config struct {
	Next     string
	Previous string
}

type locationData struct {
	Next     string              `json:"next"`
	Previous string              `json:"previous"`
	Results  []map[string]string `json:"results"`
}

func commandExit(cache *pokecache.Cache, cf *config, unused string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cache *pokecache.Cache, cf *config, unused string) error {
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

func fetchLocations(cache *pokecache.Cache, url string) (locationData, error) {
	data, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return locationData{}, fmt.Errorf("http response error: %v", err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return locationData{}, fmt.Errorf("io Read error: %v", err)
		}

		if res.StatusCode > 299 {
			return locationData{}, fmt.Errorf("response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
		}

		cache.Add(url, data)
	}

	loc := locationData{}
	if err := json.Unmarshal(data, &loc); err != nil {
		return locationData{}, fmt.Errorf("Unmarshal error: %v", err)
	}

	return loc, nil
}

func commandMap(cache *pokecache.Cache, cf *config, unused string) error {
	if cf.Next == "" {
		cf.Next = "https://pokeapi.co/api/v2/location-area/"
	}

	loc, err := fetchLocations(cache, cf.Next)
	if err != nil {
		return err
	}

	cf.Next = loc.Next
	cf.Previous = loc.Previous

	for i := range loc.Results {
		fmt.Println(loc.Results[i]["name"])
	}
	return nil
}

func commandMapb(cache *pokecache.Cache, cf *config, unused string) error {
	if cf.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	loc, err := fetchLocations(cache, cf.Previous)
	if err != nil {
		return err
	}

	cf.Next = loc.Next
	cf.Previous = loc.Previous

	for i := range loc.Results {
		fmt.Println(loc.Results[i]["name"])
	}
	return nil
}
