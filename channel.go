package main

import (
	"github.com/gorilla/websocket"

	"net/http"
	"log"
)

type appChannel struct {
	Name string
	Forward chan *Message
	Join chan *client
	Leave chan *client
	Stop chan int
	clients map[*client]bool
}

func (c *appChannel) Run() {
	for {
		select {
		case client := <-c.Join:
			c.clients[client] = true
			log.Printf("Client joined %s channel\n", c.Name)
		case client := <-c.Leave:
			delete(c.clients, client)
			//close(client.send)
			log.Printf("Client left %s channel\n", c.Name)
		case msg := <-c.Forward:
			log.Println("Forwarding message")
			log.Println(c.clients)
			for client := range c.clients {
				select {
				case client.send <- msg:
				default:
					delete(c.clients, client)
					close(client.send)
				}
			}
		//case <-c.Stop:
		//	close(c.Join)
		//	close(c.Leave)
		//	close(c.Stop)
		//	close(c.Forward)
		//	return
		}
		//for client := range c.clients {
		//	delete(c.clients, client)
		//}
	}
}

const socketBufferSize  = 1024

var upgrader = &websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
	ReadBufferSize:socketBufferSize,
	WriteBufferSize: socketBufferSize}


func newChannel(name string) *appChannel {
	return &appChannel{
		Name: 	 name,
		Forward: make(chan *Message),
		Join:    make(chan *client),
		Leave:   make(chan *client),
		Stop:   make(chan int),
		clients: make(map[*client]bool),
	}
}
