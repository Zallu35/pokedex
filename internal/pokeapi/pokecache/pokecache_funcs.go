package pokecache

import (
	"time"
)

func (c *Cache) Add(key string, value []byte) {
	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = newEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.entries[key]
	if ok {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	clock := time.NewTicker(interval)
	for t := range clock.C {
		c.mu.Lock()
		for entry := range c.entries {
			if t.Sub(c.entries[entry].createdAt) > 5 {
				delete(c.entries, entry)
			}
		}
		c.mu.Unlock()
	}

}

func NewCache(interval time.Duration) Cache {
	var activeCache Cache
	go activeCache.reapLoop(interval)
	return activeCache
}
