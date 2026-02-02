package main

import (
	"fmt"
	"pokedex/internal/pokeapi/pokecache"
)

type dex struct {
	pokemonList map[string]pokemonData
}

func commandInspect(cache *pokecache.Cache, cf *config, pkDex *dex, pokemonName string) error {
	if pokeData, ok := pkDex.pokemonList[pokemonName]; !ok {
		return fmt.Errorf("you have not caught that pokemon yet")
	} else {
		fmt.Printf("Height: %d\nWeight: %d\n", pokeData.Height, pokeData.Weight)
		fmt.Println("Base Stats:")
		for s := range len(pokeData.Stats) {
			fmt.Printf("  -%s: %d\n", pokeData.Stats[s].Stat.Name, pokeData.Stats[s].BaseStat)
		}
		fmt.Println("Types:")
		for i := range len(pokeData.Types) {
			fmt.Printf("  -%s\n", pokeData.Types[i].Type.Name)
		}
	}
	return nil
}

func commandPokedex(cache *pokecache.Cache, cf *config, pkDex *dex, unused string) error {
	if len(pkDex.pokemonList) < 1 {
		return fmt.Errorf("no pokemon have been caught yet!")
	} else {
		fmt.Println("Current Pokedex Entries:")
		for pkmn, _ := range pkDex.pokemonList {
			fmt.Printf("  -%s\n", pkmn)
		}
	}
	return nil
}
