
package main

import (
    "fmt"
    "sync"
    "time"
)

var wg sync.WaitGroup

func main() {
    start := time.Now().UTC()
    for n := 2; n <= 12; n++ {
        wg.Add(1)
        go timestable(n)
    }
    wg.Wait()

    duration := time.Now().UTC().Sub(start)
    fmt.Printf("completed in %s\n", duration)
}

func timestable(x int) {
    for i := 1; i <= 12; i++ {
        fmt.Printf("%d x %d = %d\n", x, i, x * i)
    }
    wg.Done()
}
