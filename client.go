package main

import (
	"github.com/gorilla/websocket"
	"log"
	"encoding/json"
)

type client struct {
	id string
	conn *websocket.Conn
	channel *appChannel
	send chan *Message
	incoming chan *Message
}

const messageBufferSize = 256

func (c *client) read() {
	defer c.conn.Close()
	defer func() { c.channel.Leave <- c }()
	for {
		msg := new(Message)
		if err := c.conn.ReadJSON(msg); err == nil {
			log.Println(msg)
			msg.Client = c
			c.incoming<- msg
		} else if _, ok := err.(*json.SyntaxError); ok {
			log.Println("Syntax err:", err)
		} else {
			log.Println("Socket read err:", err)
			break
		}
	}
}

func (c *client) write() {
	defer c.conn.Close()
	for msg := range c.send {
		if err := c.conn.WriteJSON(msg); err != nil {
			log.Println("Socket write err:", err)
			break
		}
	}
}

func (c *client) Run() {
	go c.write()
	c.read()
}


func newClient(conn *websocket.Conn, ch *appChannel, incomingChan chan *Message) *client {
	return &client{
		conn: 	  conn,
		send:     make(chan *Message, messageBufferSize),
		channel:  ch,
		incoming: incomingChan,
	}
}
