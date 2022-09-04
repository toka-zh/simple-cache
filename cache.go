package simple_cache

import (
	"errors"
	"sync"
)

type Cache struct {
	*cache
}

func NewCache() Cache {
	c := Cache{}
	c.cache = &cache{}
	c.cache.data = make(map[string]item)
	return c
}

type cache struct {
	data map[string]item
	mu   sync.RWMutex
}

type item struct {
	object interface{}
}

func (c *cache) Set(key string, data interface{}) error {
	if !c.mu.TryLock() {
		return errors.New("")
	}
	//c.mu.Lock()
	c.set(key, data)
	c.mu.Unlock()
	return nil
}

func (c *cache) Add(key string, data interface{}) error {
	if !c.mu.TryLock() {
		return errors.New("")
	}

	//c.mu.Lock()
	if c.exist(key) {
		c.mu.Unlock()
		return errors.New("")
	}
	c.set(key, data)
	c.mu.Unlock()
	return nil
}

func (c *cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	value, ok := c.data[key]
	c.mu.RUnlock()
	if !ok {
		return nil, false
	}
	return value.object, true
}

func (c *cache) Delete(key string) {
	if !c.mu.TryLock() {
		return
	}

	delete(c.data, key)
	c.mu.Unlock()

}

func (c *cache) set(key string, data interface{}) {
	c.data[key] = item{object: data}
}

func (c *cache) exist(key string) bool {
	if _, ok := c.data[key]; ok {
		return true
	}
	return false
}
