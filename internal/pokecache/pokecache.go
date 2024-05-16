package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]CacheEntry
	mux   *sync.Mutex
}

type CacheEntry struct {
	CreatedAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]CacheEntry),
		mux:   &sync.Mutex{},
	}

	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, vale []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache[key] = CacheEntry{
		CreatedAt: time.Now().UTC(),
		val:       vale,
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	val, ok := c.cache[key]

	return val.val, ok
}
