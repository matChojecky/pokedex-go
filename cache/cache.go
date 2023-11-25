package cache

import (
	"sync"
	"time"
)

type Cache interface {
	Add(key string, val []byte)
	Get(key string) ([]byte, bool)
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type CacheStore struct {
	store  map[string]cacheEntry
	mu *sync.Mutex
}


func NewCache(interval time.Duration) CacheStore {
	mu := sync.Mutex{}
	cache :=  CacheStore {
		store: make(map[string]cacheEntry),
		mu: &mu,
	}

	go cache.reapLoop(interval)
	
	return cache
}


func (c *CacheStore) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.store[key] = cacheEntry{
		val: val,
		createdAt: time.Now(),
	}
}

func (c CacheStore) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.store[key]

	if !ok {
		return []byte{}, false
	}

	return entry.val, true
}

func (c *CacheStore) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *CacheStore) reap(now time.Time, interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, entry := range c.store {
		if entry.createdAt.Before(now.Add(-interval)) {
			delete(c.store, key)
		}
	}

}