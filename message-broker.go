package main

import (
	"time"
	"fmt"
	"log"
)

type Message struct {
	Client *client
	Channel string `json:"channel"`
	Type string `json:"type"`
	Value string `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

type MessageBroker struct {
	Incoming chan *Message
}

const (
	message = "msg"
	join = "join"
)

func (mb *MessageBroker) Run() {
	for{
		select{
		case msg := <-mb.Incoming:
			log.Println("Incoming message")
			switch msg.Type {
			case message:
				if channel, ok := cm.Channels[msg.Channel]; ok {
					channel.Forward<- msg
				} else {
					fmt.Errorf("MessageBroker: %s channel does not exist", msg.Channel)
				}
			case join:
				if channel, ok := cm.Channels[msg.Channel]; ok {
					msg.Client.channel.Leave<- msg.Client
					channel.Join<- msg.Client
					msg.Client.send<- &Message{
						Channel: msg.Channel,
						Type: join,
						Value: "success",
					}
				} else {
					fmt.Errorf("MessageBroker: %s channel does not exist", msg.Channel)
				}
			}

		}
	}
}

func NewMessageBroker() *MessageBroker {
	return &MessageBroker{
		Incoming: make(chan *Message, messageBufferSize),
	}
}
