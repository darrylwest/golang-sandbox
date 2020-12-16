package main

import (
    "bytes"
    "fmt"
    "log"
    "time"

    "github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
        panic(err)
    }
}

func workit() {
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    failOnError(err, "failed connection")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "failed channel open")
    defer ch.Close()

    q, err := ch.QueueDeclare(
        "hello-task", // name
        true,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    failOnError(err, "failed declare queue")

    err = ch.Qos(
        1,     // prefetch
        0,     // prefetch size
        false, // global
    )
    failOnError(err, "failed to set qos")

    msgs, err := ch.Consume(
        q.Name, // routing key
        "",     // consumer
        false,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )

    failOnError(err, "failed to register a consumer")

    forever := make(chan bool)

    go func() {
        for d := range msgs {
            fmt.Printf("received: %s, please wait...", d.Body)
            count := bytes.Count(d.Body, []byte("."))
            t := time.Duration(count)
            time.Sleep(t * time.Second)
            fmt.Printf("task completed in %d seconds.\n", count)
            d.Ack(false)
        }
    }()

    fmt.Println(" [*] Waiting for messages...")
    <-forever
}

func main() {
    workit()
}

