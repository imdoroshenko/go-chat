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
	send chan message
}

func (c *client) read() {
	defer c.conn.Close()
	for {
		msg := new(message)
		if err := c.conn.ReadJSON(msg); err == nil {
			log.Println(msg)
			c.channel.forward <- *msg
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
