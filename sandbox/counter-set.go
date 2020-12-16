package main

import (
    "fmt"
    "sync"
)

type Set struct {
    sync.RWMutex
    set map[string]int64
}

func NewSet() *Set {
    return &Set{set: make(map[string]int64)}
}

func (set *Set) Add(key string) int64 {
    set.RLock()
    defer set.RUnlock()
    i, _ := set.set[key]
    i = i + 1
    set.set[key] = i
    return i
}

func (set *Set) Get(key string) int64 {
    i := set.set[key]
    return i
}

func (set *Set) Remove(key string) int64 {
    i := set.set[key]
    delete(set.set, key)
    return i
}

func main() {
    set := NewSet()
    k1 := set.Add("flarb")

    set.Add("foo")
    set.Add("foo")
    set.Add("foo")
    k2 := set.Add("foo")

    set.Add("bar")
    k3 := set.Add("bar")

    fmt.Println(k1, k2, k3)
    fmt.Println(set)

    set.Remove("bar")
    fmt.Println(set)
}




