package main

import (
    "fmt"
    "time"
)

func main() {

    seconds := 10 * time.Second
    fmt.Printf("process will run for %d seconds then exit...\n", (seconds / 1e9))

    count := 0
    ticker := time.NewTicker(300 * time.Millisecond)
    go func() {
        for now := range ticker.C {
            count++
            fmt.Printf("%d %v\n", count, now)
        }
    }()

    time.Sleep(seconds)
    ticker.Stop()
    fmt.Println("stopped...")
}
