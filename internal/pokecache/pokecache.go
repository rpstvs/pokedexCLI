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
