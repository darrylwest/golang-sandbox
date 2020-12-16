//
// cache - read only and read/write cache interfaces and implementations; read/write is thread safe
//
// @author darryl.west
// @created 2018-01-15 07:08:36
//

package cache

// ReadOnlyCache an interface for read only cache
type ReadOnlyCache interface {
	Get(string) (interface{}, bool)
	Len() int
	Keys() []string
}

// ReadWriteCache an interface to define read-write cache
type ReadWriteCache interface {
	run(map[string]interface{})
	Get(string) (interface{}, bool)
	Put(string, interface{}) bool
	Remove(string) (interface{}, bool)
	Len() int
	Close() ReadOnlyCache
}

// PutFunc a function definition to implement cache puts
type PutFunc func(interface{}, bool) interface{}

//
// ---- the read/only implementation
//

type readOnlyCache struct {
	store  map[string]interface{}
	length int
}

// NewReadOnlyCache create and return a read only cache prepared fro the hash map
func NewReadOnlyCache(hash map[string]interface{}) ReadOnlyCache {
	store := make(map[string]interface{})

	// copy it
	for k, v := range hash {
		store[k] = v
	}

	roc := readOnlyCache{
		store:  store,
		length: len(store),
	}

	return roc
}

func (roc readOnlyCache) Get(key string) (interface{}, bool) {
	value, ok := roc.store[key]
	return value, ok
}

// Len returns the length of the read only cache
func (roc readOnlyCache) Len() int {
	return roc.length
}

// Keys returns all the keys from this read only
func (roc readOnlyCache) Keys() []string {
	keys := make([]string, 0, len(roc.store))

	for k := range roc.store {
		keys = append(keys, k)
	}

	return keys
}

//
// ---- the read/write implementation
//

type commandAction int

const (
	remove commandAction = iota
	end
	get
	put
	length
)

type command struct {
	action commandAction
	key    string
	value  interface{}
}

type result struct {
	key   string
	value interface{}
	ok    bool
}

type readWriteCache struct {
	request  chan command
	response chan result
}

// NewReadWriteCache create an empty read-write cache
func NewReadWriteCache() ReadWriteCache {
	store := make(map[string]interface{}, 100)
	return ReadWriteCacheFromMap(store)
}

// ReadWriteCacheFromMap create a read/write cache with an existing backing store
func ReadWriteCacheFromMap(store map[string]interface{}) ReadWriteCache {
	rwc := readWriteCache{
		request:  make(chan command),
		response: make(chan result),
	}

	go rwc.run(store)

	return rwc
}

func (rwc readWriteCache) run(store map[string]interface{}) {
	log.Info("start the read/write cache...")

	done := make(chan bool)

    defer func() {
        close(rwc.request)
        close(rwc.response)
    }()

	for {
		var req command

		select {
		case req = <-rwc.request:
		case <-done:
			return
		}

		log.Debug("%v", req)
		res := result{}

		switch req.action {
		case remove:
			sz := len(store)
			res.key = req.key
			res.value, res.ok = store[req.key]
			if res.ok {
				log.Debug("remove %s", req.key)
				delete(store, req.key)
			}
			log.Info("size %d, new size: %d", sz, len(store))

		case end:
			hash := make(map[string]interface{})
			for k, v := range store {
				hash[k] = v
			}

			res.value = NewReadOnlyCache(hash)
			res.ok = true
			rwc.response <- res
			done <- true
		case get:
			res.key = req.key
			res.value, res.ok = store[req.key]
		case put:
			res.key = req.key
			res.value = req.value
			_, res.ok = store[req.key]
			store[req.key] = req.value
		case length:
			res.value = len(store)
			res.ok = true
		}

		rwc.response <- res
	}

}

func (rwc readWriteCache) Get(key string) (interface{}, bool) {
	req := command{
		action: get,
		key:    key,
	}

	rwc.request <- req
	result := <-rwc.response

	return result.value, result.ok
}

func (rwc readWriteCache) Put(key string, value interface{}) bool {
	req := command{
		action: put,
		key:    key,
		value:  value,
	}

	rwc.request <- req
	result := <-rwc.response

	return result.ok
}

func (rwc readWriteCache) Remove(key string) (interface{}, bool) {
	req := command{
		action: remove,
		key:    key,
	}

	rwc.request <- req
	result := <-rwc.response

	return result.value, result.ok
}

func (rwc readWriteCache) Len() int {
	req := command{
		action: length,
	}

	rwc.request <- req
	result := <-rwc.response

	return result.value.(int)
}

func (rwc readWriteCache) Close() ReadOnlyCache {
	log.Info("close the read/write cache...")

	req := command{
		action: end,
	}

	rwc.request <- req
	result := <-rwc.response

	close(rwc.request)
	// close(rwc.response)

	return result.value.(ReadOnlyCache)
}
