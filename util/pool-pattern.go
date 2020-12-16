package main

// this creates an endless pool by implmenting pool.New... should work on a top limit plus a way to shrink the pool...

import (
    "fmt"
    "sync"
    "time"
)

// an example of how to make a pool of something...
type Worker struct {
    id string
    fn func(string)
}

var currentID = 1000

func nextID() string {
    currentID++
    return fmt.Sprintf("%x%d", time.Now().UnixNano(), currentID)
}

func noop(id string) {
    fmt.Printf("id: %s ok %v\n", id, time.Now())
}

func NewPool(total int) *sync.Pool {
    pool := sync.Pool{}

    pool.New = func() interface{} {
        return &Worker{ id:nextID(), fn: noop } 
    }

    for i := 0; i < total; i++ {
        pool.Put( pool.New() )
    }

    return &pool
}

func main() {
    pool := NewPool(4)

    // this should run all 5...
    for i := 0; i < 10; i++ {
        job := pool.Get().(*Worker)
        go func() {
            job.fn(job.id)

            pool.Put(job)
        }()
    }

    fmt.Println("done...")
    time.Sleep(1 * time.Second)

}
