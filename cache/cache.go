package cache

import (
	"sync"
	"time"
)

type CacheEntry[T any] struct {
	Data       T
	Expiration time.Time
}

type Cache[T any] struct {
	data  map[string]CacheEntry[T]
	mutex sync.Mutex
}

var cacheInstances sync.Map

func GetCacheInstance[T any]() *Cache[T] {
	instance, _ := cacheInstances.LoadOrStore("cache", &Cache[T]{data: make(map[string]CacheEntry[T])})
	return instance.(*Cache[T])
}

func (c *Cache[T]) Set(key string, value T, duration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = CacheEntry[T]{Data: value, Expiration: time.Now().Add(duration)}
}

// Get method to retrieve any object from cache
func (c *Cache[T]) Get(key string) (T, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, exists := c.data[key]
	if !exists || time.Now().After(entry.Expiration) {
		delete(c.data, key) // Remove expired entry
		var empty T         // Default zero value of T
		return empty, false
	}
	return entry.Data, true
}
