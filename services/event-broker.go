package services

import (
	"log"
	"github.com/imdoroshenko/go-chat/models"
	"runtime"
)

type EventBroker struct {
	Incoming chan *models.Event
	ChannelManager *ChannelManager
	Resources map[string] chan *models.Event
}

func (eb *EventBroker) Run() {
	for event := range eb.Incoming {
		if channel, ok := eb.Resources[event.Resource]; ok {
			log.Println("Incoming event", event)
			log.Println("Go rutienses", runtime.NumGoroutine())
			channel<- event
		} else {
			log.Println("Invalid event", event)
		}
	}
}

const messageBufferSize = 256

func NewEventBroker() *EventBroker {
	return &EventBroker{
		Incoming: make(chan *models.Event, messageBufferSize),
		Resources: make(map[string] chan *models.Event),
	}
}
