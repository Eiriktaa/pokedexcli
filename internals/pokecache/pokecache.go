package pokecache

import (
	"sync"
	"time"
)

type PokeCache struct {
	CacheMap      map[string]pokeCacheEntry
	Mutex         *sync.RWMutex
	validDuration time.Duration
	ticker        time.Ticker
}

type pokeCacheEntry struct {
	createdAt time.Time
	value     []byte
}

func NewCache(interval time.Duration) PokeCache {

	cache := PokeCache{
		CacheMap:      map[string]pokeCacheEntry{},
		Mutex:         &sync.RWMutex{},
		validDuration: interval,
		ticker:        *time.NewTicker(interval),
	}
	cache.reapLoop()
	return cache
}

func (c PokeCache) AddData(key string, data []byte) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	c.CacheMap[key] = pokeCacheEntry{createdAt: time.Now(), value: data}
}

func (c PokeCache) GetData(key string) ([]byte, bool) {
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()

	data, ok := c.CacheMap[key]
	if !ok {
		return nil, false
	}
	return data.value, true
}

// invalidates cache entries on a given interval
func (c PokeCache) reapLoop() {
	go func() {
		for {
			<-c.ticker.C
			c.Mutex.Lock()
			for k, v := range c.CacheMap {
				// cache entry expired
				if v.createdAt.Add(c.validDuration).Before(time.Now()) {
					delete(c.CacheMap, k)
				}
			}
			c.Mutex.Unlock()
		}
	}()
}
