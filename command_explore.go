package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedex/internal/pokeapi"
	"pokedex/internal/pokeapi/pokecache"
)

type locationDetails struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func commandExplore(cache *pokecache.Cache, cf *config, zone string) error {
	fullurl := pokeapi.BaseURL + "location-area/" + zone
	details, err := fetchLocationDetails(cache, fullurl)
	if err != nil {
		return fmt.Errorf("Explore failed to fetch location details: %v", err)
	}
	for _, pkmn := range details.PokemonEncounters {
		fmt.Println(pkmn.Pokemon.Name)
	}
	return nil
}

func fetchLocationDetails(cache *pokecache.Cache, url string) (locationDetails, error) {
	data, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return locationDetails{}, fmt.Errorf("http response error: %v", err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return locationDetails{}, fmt.Errorf("io Read error: %v", err)
		}

		if res.StatusCode > 299 {
			return locationDetails{}, fmt.Errorf("response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
		}

		cache.Add(url, data)
	}

	details := locationDetails{}
	if err := json.Unmarshal(data, &details); err != nil {
		return locationDetails{}, fmt.Errorf("error unmarshaling details: %v", err)
	}

	return details, nil
}
