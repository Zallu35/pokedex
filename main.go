package main

import (
	"pokedex/internal/pokeapi/pokecache"
	"time"
)

func main() {
	theCache := pokecache.NewCache(5 * time.Second)
	theDex := dex{pokemonList: make(map[string]pokemonData)}
	startREPL(theCache, &theDex)

}
