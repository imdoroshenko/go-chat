package actions

import (
	"github.com/imdoroshenko/go-chat/models"
	"github.com/imdoroshenko/go-chat/services"
	"fmt"
)


const (
	send = "send"
	join = "join"
	Message = "msg"
	Channel = "chn"
)

type ChannelHandler struct {
	Incoming chan *models.Event
	ChannelManager *services.ChannelManager
}

func (ch *ChannelHandler) Run() {
	for event := range ch.Incoming {
		switch event.Method {
		case join:
			if channel, ok := ch.ChannelManager.Channels[event.Channel.Name]; ok {
				event.Client.Channel.Leave<- event.Client
				channel.Join<- event.Client
				event.Status = 200
				event.Client.Send<- event
			} else {
				fmt.Errorf("MessageBroker: %s channel does not exist", event.Channel.Name)
			}
		}
	}
}
