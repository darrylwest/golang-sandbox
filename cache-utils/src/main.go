package main

import (
	"cache"
	"fmt"
    "time"

	"github.com/darrylwest/go-unique/unique"
)

func perfTests() {
	sz := int(1e5)

    store := make(map[string]interface{}, sz)
    rwc := cache.ReadWriteCacheFromMap(store)
    defer rwc.Close()

    keys := make([]string, 0, sz)
    for i := 0; i < sz; i++ {
        keys = append(keys, unique.CreateULID())
    }

    t0 := time.Now()
    for _, key := range keys {
        rwc.Put(key, key)
    }

    t1 := time.Now()
    dur := t1.Sub(t0)

    fmt.Printf(">> time: for %d puts = %s, %f ms per put\n", sz, dur.String(), float64(int(dur)/sz)/1000.0)
}

func main() {
    cache.CreateLogger()
    perfTests()
}
