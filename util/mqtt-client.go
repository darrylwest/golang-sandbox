package main

import (
    "fmt"
    "github.com/eclipse/paho.mqtt.golang"
    "time"
)

func main() {
    publish := false

    opts := mqtt.NewClientOptions()

    opts.AddBroker("tcp://127.0.0.1:1883")
    opts.SetClientID(fmt.Sprintf("%d", time.Now().Unix()))

    topic := "ebay/local.cache"
    qos := byte(2)

    if publish {
        client := mqtt.NewClient(opts)
        if token := client.Connect(); token.Wait() && token.Error() != nil {
            panic(token.Error())
        }

        fmt.Printf("client connected publishing to %s", topic)
        for i := 0; i < 10; i++ {
            payload := fmt.Sprintf("set %s %s %d", "mykey", "my string value:", i+100)
            fmt.Printf("topic: %s msg: %s\n", topic, payload)
            token := client.Publish(topic, qos, false, payload)
            token.Wait()
        }

        time.Sleep(time.Second)
    } else {
        receivedCount := 0
        choke := make(chan [2]string)

        opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
            choke <- [2]string{msg.Topic(), string(msg.Payload())}
        })

        client := mqtt.NewClient(opts)
        if token := client.Connect(); token.Wait() && token.Error() != nil {
            panic(token.Error())
        }

        if token := client.Subscribe(topic, qos, nil); token.Wait() && token.Error() != nil {
            panic(token.Error())
        }

        defer client.Disconnect(250)

        for {
            incoming := <- choke
            receivedCount++
            fmt.Printf("rcvd %d: %s %s\n", receivedCount, incoming[0], incoming[1])
        }

    }
}

