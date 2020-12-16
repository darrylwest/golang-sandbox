package main

// This is an old version that uses locks.  a much better version is available that uses channels and has an interface; see ../golang/cache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	values map[string]interface{}
	sync.RWMutex
}

func NewCache() *Cache {
	cache := new(Cache)

	cache.values = make(map[string]interface{})

	return cache
}

func (c *Cache) Len() int {
	return len(c.values)
}

func (c *Cache) Set(key string, value interface{}) error {
	fmt.Println("set:", key, value)
	c.Lock()
	c.values[key] = value
	c.Unlock()

	return nil
}

func (c *Cache) Get(key string) (interface{}, error) {
	fmt.Println("get:", key)
	c.Lock()
	value := c.values[key]
	c.Unlock()

	return value, nil
}

func (c *Cache) Delete(key string) interface{} {
	fmt.Println("delete:", key)

	if value, ok := c.values[key]; ok {
		c.Lock()
		delete(c.values, key)
		c.Unlock()
		return value
	} else {
		return nil
	}
}

func main() {

	cache := NewCache()

	key := "mykey"
	cache.Set(key, "my value")

	v, err := cache.Get(key)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("key %s returned value: %v\n", key, v)

	key = "unknown"
	v, err = cache.Get(key)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("returned for key %s, value: %v\n", key, v)

	for i := 0; i < 10; i++ {
		cache.Set(fmt.Sprintf("key-%d", i+101), fmt.Sprintf("value %v", time.Now()))
	}

	fmt.Println("cache size:", cache.Len())

	key = "mykey"
	v = cache.Delete(key)

	fmt.Println("deleted:", key, "=", v)

	fmt.Println("cache size:", cache.Len())

}
