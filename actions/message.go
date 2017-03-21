package actions

import (
	"github.com/imdoroshenko/go-chat/models"
	"github.com/imdoroshenko/go-chat/services"
	"log"
)

type MessageHandler struct {
	Incoming chan *models.Event
	ChannelManager *services.ChannelManager
}

func (mh *MessageHandler) Run() {
	for event := range mh.Incoming {
		switch event.Method {
		case send:
			if channel, ok := mh.ChannelManager.Channels[event.Message.Channel]; ok {
				event.Status = 200
				channel.Send<- event
			} else {
				log.Printf("MessageBroker: %s channel does not exist\n", event.Message.Channel)
			}
		}
	}
}
