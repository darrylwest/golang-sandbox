package main

import (
    "fmt"
    "log"

    "github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
        panic(err)
    }
}

func receiveit() {
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    failOnError(err, "failed connection")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "failed channel open")
    defer ch.Close()

    q, err := ch.QueueDeclare(
        "hello", // name
        false,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    failOnError(err, "failed declare queue")

    msgs, err := ch.Consume(
        q.Name, // routing key
        "",     // consumer
        true,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )

    failOnError(err, "failed to register a consumer")

    forever := make(chan bool)

    go func() {
        for d := range msgs {
            fmt.Printf("received: %s\n", d.Body)
        }
    }()

    fmt.Println(" [*] Waiting for messages...")
    <-forever
}

func main() {
    receiveit()
}

