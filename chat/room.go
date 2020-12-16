package main

import (
	"../trace"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)

type room struct {
	forward chan []byte
	join    chan *client
	leave   chan *client
	clients map[*client]bool
	tracer  trace.Tracer
}

func NewRoom() *room {
	log.Println("create and return the room")
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		tracer:  trace.New(os.Stdout),
	}
}

func (r *room) run() {
	r.tracer.Trace("run the room, clients: ", r.clients)
	for {
		r.tracer.Trace("for loop")
		select {
		case client := <-r.join:
			r.tracer.Trace("add a client")
			r.clients[client] = true
		case client := <-r.leave:
			r.tracer.Trace("remove a client")
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			r.tracer.Trace("process a message")
			for client := range r.clients {
				select {
				case client.send <- msg:
					// send the message
				default:
					r.tracer.Trace("got default, leaving...")
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
	r.tracer.Trace("leave the room")
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.tracer.Trace("create the upgrade")
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("serve http: ", err)
		return
	}

	r.tracer.Trace("serving...")

	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}

	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}
