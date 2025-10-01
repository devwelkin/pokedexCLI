package pokecache

import (
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	cache := NewCache(5 * time.Second)
	if cache == nil {
		t.Fatal("NewCache returned nil")
	}
	if cache.entries == nil {
		t.Error("cache.entries is nil")
	}
	if cache.stop == nil {
		t.Error("cache.stop is nil")
	}
	cache.Stop()
}

func TestCacheAddAndGet(t *testing.T) {
	cache := NewCache(5 * time.Second)
	defer cache.Stop()

	key := "test-key"
	value := []byte("test-value")

	// Test Add
	cache.Add(key, value)

	// Test Get
	retrieved, found := cache.Get(key)
	if !found {
		t.Error("Expected to find key in cache")
	}
	if string(retrieved) != string(value) {
		t.Errorf("Expected %s, got %s", string(value), string(retrieved))
	}
}

func TestCacheGetNonExistent(t *testing.T) {
	cache := NewCache(5 * time.Second)
	defer cache.Stop()

	_, found := cache.Get("non-existent-key")
	if found {
		t.Error("Expected not to find non-existent key")
	}
}

func TestCacheExpiration(t *testing.T) {
	interval := 100 * time.Millisecond
	cache := NewCache(interval)
	defer cache.Stop()

	key := "expiring-key"
	value := []byte("expiring-value")

	cache.Add(key, value)

	// Should be found immediately
	_, found := cache.Get(key)
	if !found {
		t.Error("Expected to find key immediately after adding")
	}

	// Wait for expiration
	time.Sleep(interval + 50*time.Millisecond)

	// Should not be found after expiration
	_, found = cache.Get(key)
	if found {
		t.Error("Expected key to be expired")
	}
}

func TestCacheReapLoop(t *testing.T) {
	interval := 100 * time.Millisecond
	cache := NewCache(interval)
	defer cache.Stop()

	// Add multiple entries
	for i := 0; i < 5; i++ {
		cache.Add(string(rune('a'+i)), []byte{byte(i)})
	}

	// Wait for reap loop to run
	time.Sleep(interval + 150*time.Millisecond)

	// Verify entries are cleaned up
	cache.mutex.RLock()
	entryCount := len(cache.entries)
	cache.mutex.RUnlock()

	if entryCount != 0 {
		t.Errorf("Expected 0 entries after reap, got %d", entryCount)
	}
}

func TestCacheStop(t *testing.T) {
	cache := NewCache(1 * time.Second)

	// Stop the cache
	cache.Stop()

	// Try to add after stop (should not panic)
	cache.Add("key", []byte("value"))

	// Get should still work
	_, _ = cache.Get("key")
}

func TestCacheConcurrency(t *testing.T) {
	cache := NewCache(5 * time.Second)
	defer cache.Stop()

	done := make(chan bool)
	
	// Concurrent writes
	for i := 0; i < 10; i++ {
		go func(n int) {
			cache.Add(string(rune('a'+n)), []byte{byte(n)})
			done <- true
		}(i)
	}

	// Wait for all writes
	for i := 0; i < 10; i++ {
		<-done
	}

	// Concurrent reads
	for i := 0; i < 10; i++ {
		go func(n int) {
			cache.Get(string(rune('a' + n)))
			done <- true
		}(i)
	}

	// Wait for all reads
	for i := 0; i < 10; i++ {
		<-done
	}
}
