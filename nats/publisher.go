package main

import (
    "fmt"
    "github.com/nats-io/go-nats"
    "time"
)

func main() {
    nc, err := nats.Connect(nats.DefaultURL)
    if err != nil {
        panic(err)
    }

    topic := "repl"
    for i := 0; i < 100; i++ {
        msg := fmt.Sprintf("data message: %d", time.Now().Unix())
        nc.Publish(topic, []byte(msg))
        fmt.Printf("published: %s\n", msg)
        // time.Sleep(time.Second)
    }
}
