package main

import (
	"github.com/gorilla/websocket"

	"net/http"
	"log"
)

type appChannel struct {
	name string
	forward chan message
	join chan *client
	leave chan *client
	clients map[*client]bool
}


func (c *appChannel) run() {
	for {
		select {
		case client := <-c.join:
			c.clients[client] = true
			log.Printf("Client joined %s channel\n", c.name)
		case client := <-c.leave:
			delete(c.clients, client)
			close(client.send)
			log.Printf("Client left %s channel\n", c.name)
		case msg := <-c.forward:
			for client := range c.clients {
				select {
				case client.send <- msg:
				default:
					delete(c.clients, client)
					close(client.send)
				}
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
	ReadBufferSize:socketBufferSize,
	WriteBufferSize: socketBufferSize}

func (c *appChannel) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		conn: 	 conn,
		send:    make(chan message, messageBufferSize),
		channel: c,
	}
	c.join <- client
	defer func() { c.leave <- client }()
	go client.write()
	client.read()
}



func newChannel(name string) *appChannel {
	return &appChannel{
		name: 	 name,
		forward: make(chan message),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}
