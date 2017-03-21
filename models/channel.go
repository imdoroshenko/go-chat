package models

import (
	"log"
)

type Channel struct {
	Name string `json:"name"`
	Send chan *Event `json:"-"`
	Join chan *Client `json:"-"`
	Leave chan *Client `json:"-"`
	Stop chan int `json:"-"`
	clients map[*Client]bool `json:"-"`
}

func (c *Channel) Run() {
	for {
		select {
		case client := <-c.Join:
			c.clients[client] = true
			client.Channel = c
			log.Printf("Client joined %s channel\n", c.Name)
		case client := <-c.Leave:
			delete(c.clients, client)
			//close(client.send)
			log.Printf("Client left %s channel\n", c.Name)
		case e := <-c.Send:
			log.Println("Forwarding event")
			log.Println(c.clients)
			for client := range c.clients {
				select {
				case client.Send<- e:
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
		Name:	name,
		Send: make(chan *Event),
		Join:	make(chan *Client),
		Leave:	make(chan *Client),
		Stop:	make(chan int),
		clients:	make(map[*Client]bool),
	}
}
