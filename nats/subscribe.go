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

    timeout := time.Second * 60
    topic := "repl"
    sub, err := nc.SubscribeSync(topic)
    for {
        fmt.Println("listen for message...")
        m, err := sub.NextMsg(timeout)
        if err != nil {
            panic(err)
        }
        fmt.Printf("topic: %s : %s\n", m.Subject, m.Data)
    }
}
