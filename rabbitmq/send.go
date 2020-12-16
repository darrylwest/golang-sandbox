package main

import (
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

func sendit() {
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

    body := fmt.Sprintf("hello rabbit mq! @ %v", time.Now())
    err = ch.Publish(
        "",     // exchange
        q.Name, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing {
            ContentType: "text/plain",
            Body:   []byte(body),
        })

    fmt.Printf(" [x] Sent %s\n", body)
    failOnError(err, "publish failed")
}

func main() {
    sendit()
}

