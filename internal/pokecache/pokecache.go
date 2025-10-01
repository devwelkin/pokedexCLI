package pokecache

import (
	"sync"
	"time"
)

// cacheEntry represents a single entry in the cache
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries  map[string]cacheEntry
	mutex    sync.RWMutex
	interval time.Duration
	stop     chan struct{}
}

// NewCache creates a new cache instance with a configurable interval
func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
		stop:     make(chan struct{}),
	}

	// Start the cleanup process in the background
	go cache.reapLoop()

	return cache
}

// reapLoop periodically cleans up expired entries
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.mutex.Lock()
			for key, entry := range c.entries {
				// Delete entries older than the specified interval
				if time.Since(entry.createdAt) > c.interval {
					delete(c.entries, key)
				}
			}
			c.mutex.Unlock()
		case <-c.stop:
			return // Stop the cleaning
		}
	}
}

// Add adds a new entry to the cache
func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// Get retrieves an entry from the cache
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	entry, exists := c.entries[key]
	if !exists {
		return nil, false
	}
	// Check if entry has expired
	if time.Since(entry.createdAt) > c.interval {
		return nil, false
	}
	return entry.val, true
}

// Stop stops the cache cleanup goroutine
func (c *Cache) Stop() {
	close(c.stop)
}
