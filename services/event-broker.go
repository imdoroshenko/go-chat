package services

import (
	"fmt"
	"log"
	"github.com/imdoroshenko/go-chat/models"
)

type EventBroker struct {
	Incoming chan *models.Event
	ChannelManager *ChannelManager
}

const (
	message = "msg"
	join = "join"
)

func (eb *EventBroker) Run() {
	for{
		select{
		case msg := <-eb.Incoming:
			log.Println("Incoming message")
			switch msg.Type {
			case message:
				if channel, ok := eb.ChannelManager.Channels[msg.Channel]; ok {
					channel.Forward<- msg
				} else {
					fmt.Errorf("MessageBroker: %s channel does not exist", msg.Channel)
				}
			case join:
				if channel, ok := eb.ChannelManager.Channels[msg.Channel]; ok {
					msg.Client.Channel.Leave<- msg.Client
					channel.Join<- msg.Client
					msg.Client.Send<- &models.Event{
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

func NewEventBroker() *EventBroker {
	return &EventBroker{
		Incoming: make(chan *models.Event, messageBufferSize),
	}
}
