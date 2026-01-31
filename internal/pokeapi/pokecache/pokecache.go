package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	activeCache := Cache{entries: make(map[string]cacheEntry)}
	go activeCache.reapLoop(interval)
	return &activeCache
}
