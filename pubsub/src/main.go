package main

import (
    "fmt"
    "pubsub"
    "time"
)

func main() {
    fmt.Println("pub-sub")

    broker := pubsub.NewBroker()
    todoSub,err := broker.Attach()
    if err != nil {
        panic(err)
    }
    eventSub,err := broker.Attach()
    if err != nil {
        panic(err)
    }

    broker.Subscribe(todoSub, "todo")
    broker.Subscribe(eventSub, "event")

    go func() {
        ch := todoSub.GetMessages()
        fmt.Printf("todo: %T\n", ch)
        for {
            msg := <- ch
            t := msg.GetCreatedAt()
            sec := t / 1e9
            nano := t % 1e9
            dt := time.Unix(sec, nano)

            fmt.Printf("todo: %v, %s\n", msg.GetPayload(), dt.Format(time.RFC3339))
        }
    }()

    go func() {
        ch := eventSub.GetMessages()
        fmt.Printf("event: %T\n", ch)
        for {
            msg := <- ch
            fmt.Printf("event: %v, %d\n", msg.GetPayload(), msg.GetCreatedAt())
        }
    }()

    broker.Broadcast("my todo list", "todo")
    broker.Broadcast("my new event", "event")

    time.Sleep(2 * time.Second)
    broker.Broadcast("my second todo list", "todo")
    broker.Broadcast("my second event", "event")

    time.Sleep(5 * time.Second)
}

