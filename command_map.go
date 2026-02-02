package main

import (
	"encoding/json"
	"fmt"
	"pokedex/internal/pokeapi/pokecache"
)

type locationData struct {
	Next     string              `json:"next"`
	Previous string              `json:"previous"`
	Results  []map[string]string `json:"results"`
}

func fetchLocations(cache *pokecache.Cache, url string) (locationData, error) {
	data, err := fetchData(cache, url)
	if err != nil {
		return locationData{}, fmt.Errorf("fetchData failed with error: %s", err)
	}

	loc := locationData{}
	if err := json.Unmarshal(data, &loc); err != nil {
		return locationData{}, fmt.Errorf("Unmarshal error: %v", err)
	}

	return loc, nil
}

func commandMap(cache *pokecache.Cache, cf *config, a *dex, unused string) error {
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

func commandMapb(cache *pokecache.Cache, cf *config, a *dex, unused string) error {
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
