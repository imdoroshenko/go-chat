package models

import (
	"github.com/gorilla/websocket"
	"log"
	"encoding/json"
)

type Client struct {
	id string
	conn *websocket.Conn
	Channel *Channel
	Send chan *Message
	incoming chan *Message
}

const messageBufferSize = 256

func (c *Client) read() {
	defer c.conn.Close()
	defer func() { c.Channel.Leave <- c }()
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

func (c *Client) write() {
	defer c.conn.Close()
	for msg := range c.Send {
		if err := c.conn.WriteJSON(msg); err != nil {
			log.Println("Socket write err:", err)
			break
		}
	}
}

func (c *Client) Run() {
	go c.write()
	c.read()
}


func NewClient(conn *websocket.Conn, ch *Channel, incomingChan chan *Message) *Client {
	return &Client{
		conn: 	  conn,
		Send:     make(chan *Message, messageBufferSize),
		Channel:  ch,
		incoming: incomingChan,
	}
}
