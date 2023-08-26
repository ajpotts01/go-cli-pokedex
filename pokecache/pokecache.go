package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	lock  *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		cache: make(map[string]cacheEntry),
		lock:  &sync.Mutex{},
	}

	go newCache.reapLoop(interval)
	return newCache
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for {
		t := <-ticker.C

		for key, entry := range c.cache {
			if t.Sub(entry.createdAt).Milliseconds() >= interval.Milliseconds() {
				if _, ok := c.cache[key]; ok {
					c.lock.Lock()
					delete(c.cache, key)
					c.lock.Unlock()
				}
			}
		}
	}
}

func (c Cache) Add(key string, val []byte) {
	var entry cacheEntry

	entry.createdAt = time.Now()
	entry.val = val

	c.lock.Lock()
	defer c.lock.Unlock()
	c.cache[key] = entry
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	_, ok := c.cache[key]

	if !ok {
		return nil, false
	}

	return c.cache[key].val, true
}
