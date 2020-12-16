//
// cache tests
//
// @author darryl.west@ebay.com
// @created 2017-11-27 12:57:59
//

package unit

import (
	"cache"
	"fmt"

	"testing"
	"time"

	"github.com/darrylwest/go-unique/unique"

	. "github.com/franela/goblin"
)

func createTestHash() map[string]interface{} {
	hash := make(map[string]interface{})

	hash["flarb"] = "one original flarber"
	hash["flub"] = "two original flubber"

	return hash
}

func TestCache(t *testing.T) {
	g := Goblin(t)

	g.Describe("Cache", func() {
		log := cache.CreateLogger()
		log.SetLevel(3)

		g.It("should create a read only cache struct", func() {
			hash := createTestHash()
			roc := cache.NewReadOnlyCache(hash)
			g.Assert(fmt.Sprintf("%T", roc)).Equal("cache.readOnlyCache")
			g.Assert(roc.Len()).Equal(len(hash))
		})

		g.It("should return a known set of keys from the read only cache", func() {
			hash := createTestHash()
			roc := cache.NewReadOnlyCache(hash)

			keys := roc.Keys()
			log.Debug("keys: %v", keys)

			g.Assert(len(keys)).Equal(len(hash))
			for _, key := range keys {
				val, ok := roc.Get(key)
				g.Assert(ok).IsTrue()

				g.Assert(val).Equal(hash[key])
			}
		})

		g.It("should retain values even if original hash changes", func() {
			hash := createTestHash()
			roc := cache.NewReadOnlyCache(hash)

			// now change the underlying map and verify read-only...
			hash["flarb"] = "my new flarb"
			hash["flub"] = "a new flub"
			hash["flark"] = "a new element..."

			log.Debug("hash: %v", hash)

			g.Assert(roc.Len()).Equal(len(hash) - 1)
			for _, key := range roc.Keys() {
				val, ok := roc.Get(key)
				g.Assert(ok).IsTrue()

				g.Assert(val != hash[key]).IsTrue()
			}

			if val, ok := roc.Get("flark"); ok {
				g.Assert(val).Equal(nil)
			}
		})

		g.It("should create a new a read write cache", func() {
			rwc := cache.NewReadWriteCache()
			defer rwc.Close()

			g.Assert(fmt.Sprintf("%T", rwc)).Equal("cache.readWriteCache")
			g.Assert(rwc.Len()).Equal(0)
		})

		g.It("should get/put values for a read write cache", func() {
			key := "mykey"
			value := "this is a test value"

			rwc := cache.NewReadWriteCache()
			defer rwc.Close()

			ok := rwc.Put(key, value)
			g.Assert(ok).Equal(false)

			v, ok := rwc.Get(key)
			g.Assert(ok).Equal(true)
			g.Assert(v).Equal(value)
		})

		g.It("should remove a value from a read write cache", func() {
			hash := createTestHash()
			rwc := cache.NewReadWriteCache()
			defer rwc.Close()

			for k, v := range hash {
				rwc.Put(k, v)
			}

			g.Assert(rwc.Len()).Equal(len(hash))

			key := "flarb"
			v, ok := rwc.Remove(key)
			g.Assert(ok).Equal(true)
			g.Assert(v).Equal(hash[key])

			g.Assert(rwc.Len()).Equal(len(hash) - 1)
		})

		g.It("should close the current read-write cache and return a read only cache", func() {
			hash := createTestHash()
			rwc := cache.NewReadWriteCache()

			for k, v := range hash {
				rwc.Put(k, v)
			}

			g.Assert(rwc.Len()).Equal(len(hash))

			roc := rwc.Close()
			g.Assert(fmt.Sprintf("%T", roc)).Equal("cache.readOnlyCache")
			g.Assert(roc.Len()).Equal(len(hash))

			// TODO handle attempted access after close...
			// rwc.Put("testkey", "my test value")
		})

		g.It("should accept a large number of puts in a resonable amount of time", func() {
			sz := 100000

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

			g.Assert(rwc.Len()).Equal(sz)
			g.Assert(dur < time.Duration(time.Millisecond*200)).IsTrue()
		})

		g.It("should create a new a read write cache from a non-empty hash")
	})
}
