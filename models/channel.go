package models

import (
	"log"
)

type Channel struct {
	Name string
	Forward chan *Message
	Join chan *Client
	Leave chan *Client
	Stop chan int
	clients map[*Client]bool
}

func (c *Channel) Run() {
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
				case client.Send <- msg:
				default:
					delete(c.clients, client)
					close(client.Send)
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


func NewChannel(name string) *Channel {
	return &Channel{
		Name: 	 name,
		Forward: make(chan *Message),
		Join:    make(chan *Client),
		Leave:   make(chan *Client),
		Stop:   make(chan int),
		clients: make(map[*Client]bool),
	}
}
