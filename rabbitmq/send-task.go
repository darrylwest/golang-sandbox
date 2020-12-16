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

func sendit(task string) {
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

    body := fmt.Sprintf("%x-%s", time.Now().Unix(), task)
    err = ch.Publish(
        "",     // exchange
        q.Name, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing {
            DeliveryMode: amqp.Persistent,
            ContentType: "text/plain",
            Body:   []byte(body),
        })

    fmt.Printf(" [x] Sent %s\n", body)
    failOnError(err, "publish failed")
}

func main() {
    tasks := []string{ "t1...", "t2....", "t3..", "t4...", "t5....." }
    for _, task := range tasks {
        sendit(task)
    }
}

