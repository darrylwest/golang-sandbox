// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// implementation of a JSON request client

package main

import (
	"flag"
    "fmt"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/wsapi"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	go func() {
		defer c.Close()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Millisecond * 100)
	defer ticker.Stop()

    count := 0
	for range ticker.C {
		// err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))

        count++
        msg := make(map[string]interface{})
        msg["ID"] = fmt.Sprintf("%d-%d", time.Now().Unix(), count)
        msg["Tests"] = 24
        msg["Failed"] = 0

        err := c.WriteJSON(msg)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
