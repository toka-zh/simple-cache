package test

import (
	"fmt"
	"simple-cache"
	"testing"
)

func TestSimpleCache(t *testing.T) {
	cache := simple_cache.NewCache()
	err := cache.Add("testKey1", "testValue1")
	if err != nil {
		return
	}
	fmt.Println(cache.Get("testKey1"))
	cache.Delete("testKey1")
	fmt.Println(cache.Get("testKey1"))
}
