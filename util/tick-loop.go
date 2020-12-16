package main

import (
    "fmt"
    "time"
)

func main() {
    timechan := time.NewTimer(time.Second * 2).C
    tickchan := time.NewTicker(time.Millisecond * 400).C

    done := make(chan bool)

    go func() {
        time.Sleep(time.Second * 10)
        done <- true
    }()

    count := 0

    for {
        select {
        case <- timechan:
            fmt.Println("timer expired...")
        case <- tickchan:
            count++
            fmt.Printf("ticker ticks: %d\n", count)
        case <- done:
            fmt.Println("done now...")
            return
        }
    }
}
