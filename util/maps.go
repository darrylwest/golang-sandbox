package main

// TODO : refactor to use channels to replace mutex...

import (
	"fmt"
	"sync"
	"time"
)

type IntMap struct {
	sync.RWMutex
	hash map[string]int64
}

func NewIntMap() *IntMap {
	return &IntMap{hash: make(map[string]int64)}
}

func (m *IntMap) put(key string, value int64) int64 {
	m.RLock()
	defer m.RUnlock()

	m.hash[key] = value

	return value
}

func (m *IntMap) get(key string) (int64, bool) {
	v, ok := m.hash[key]
	return v, ok
}

func main() {
	hash := NewIntMap()
	key := "/mykey\n"
	value := time.Now().Unix()

	hash.put(key, value)

	fmt.Println(hash)

	fmt.Println("-> check for known key...")
	if v, ok := hash.get(key); ok {
		fmt.Printf("t = %d, ok: %v\n", v, ok)
	} else {
		fmt.Println("error, should have a value")
	}
}
