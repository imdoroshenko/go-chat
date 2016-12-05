package services

import (
	"fmt"
	"log"
	"github.com/imdoroshenko/go-chat/models"
)

type MessageBroker struct {
	Incoming chan *models.Message
	ChannelManager *ChannelManager
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
				if channel, ok := mb.ChannelManager.Channels[msg.Channel]; ok {
					channel.Forward<- msg
				} else {
					fmt.Errorf("MessageBroker: %s channel does not exist", msg.Channel)
				}
			case join:
				if channel, ok := mb.ChannelManager.Channels[msg.Channel]; ok {
					msg.Client.Channel.Leave<- msg.Client
					channel.Join<- msg.Client
					msg.Client.Send<- &models.Message{
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

const messageBufferSize = 256

func NewMessageBroker() *MessageBroker {
	return &MessageBroker{
		Incoming: make(chan *models.Message, messageBufferSize),
	}
}
