package main

import (
	"fmt"
	"sync"
	"time"
)

const defaultExpirationTime = 10 * time.Second

type CacheItem struct {
	Value      any
	Expiration int64
}

type Cache struct {
	mu    sync.RWMutex
	items map[string]CacheItem
}

func (c *Cache) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = CacheItem{
		Value:      value,
		Expiration: time.Now().Add(defaultExpirationTime).UnixNano(),
	}
}

func (c *Cache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, ok := c.items[key]
	if !ok {
		return nil, false
	}

	return item.Value, true
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	fmt.Println("deleting key", key)

	delete(c.items, key)
}

func (c *Cache) PrintCache() {
	c.mu.RLock()
	defer c.mu.RUnlock()

	fmt.Println("Cache items:")

	for key, item := range c.items {
		fmt.Printf("\t%s: %v", key, item)
	}
}

func (c *Cache) evict() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, item := range c.items {
		if time.Now().UnixNano() > item.Expiration {
			fmt.Println("\tEvicting: ", key)
			delete(c.items, key)
		}
	}
}

func (c *Cache) StartEvict() {
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ticker.C:
			c.evict()
		}
	}
}

func NewCache() *Cache {
	c := &Cache{
		items: make(map[string]CacheItem),
	}

	go c.StartEvict()

	return c
}

func main() {
	cache := NewCache()

	go func() {
		for i := 0; i < 10; i++ {
			cache.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
			time.Sleep(750 * time.Millisecond)
		}
	}()

	cache.Set("key1", "value")
	value, ok := cache.Get("key1")
	if ok {
		println(value.(string))
	}

	_, ok = cache.Get("key2")
	if !ok {
		println("key2 not found")
	}

	counter := 0

	for counter < 60 {
		fmt.Printf("Cache size: %d\n", len(cache.items))
		counter++

		if len(cache.items) == 0 {
			break
		}

		time.Sleep(2 * time.Second)
	}
}
