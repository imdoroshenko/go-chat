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
	Send chan *Event
	incoming chan *Event
}

const eventBufferSize = 256

func (c *Client) read() {
	defer func() {
		c.Channel.Leave <- c
		close(c.Send)
		c.conn.Close()
	}()

	for {
		e := new(Event)
		if err := c.conn.ReadJSON(e); err == nil {
			log.Println(e)
			e.Client = c
			c.incoming<- e
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
	for e := range c.Send {
		if err := c.conn.WriteJSON(e); err != nil {
			log.Println("Socket write err:", err)
			break
		}
	}
}

func (c *Client) Run() {
	go c.write()
	c.read()
}


func NewClient(conn *websocket.Conn, ch *Channel, incomingChan chan *Event) *Client {
	return &Client{
		conn: 	  conn,
		Send:     make(chan *Event, eventBufferSize),
		Channel:  ch,
		incoming: incomingChan,
	}
}
