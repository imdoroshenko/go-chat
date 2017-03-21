package main

import (
	"github.com/imdoroshenko/go-chat/services"
	"github.com/imdoroshenko/go-chat/models"
	"github.com/imdoroshenko/go-chat/actions"
)

const (
	defaultChannel = "default"
	eventBufferSize = 256
)

var (
	ChannelManager *services.ChannelManager
	EventBroker *services.EventBroker
)

func init() {
	ChannelManager = services.NewChannelManager()
	EventBroker = services.NewEventBroker()

	EventBroker.ChannelManager = ChannelManager

	go EventBroker.Run()
	go ChannelManager.Run()

	ChannelManager.Open<- defaultChannel
	ChannelManager.Open<- "other"

	// init resource handle channels
	channelEvents := make(chan *models.Event, eventBufferSize)
	messageEvents := make(chan *models.Event, eventBufferSize)

	channelsHandler := &actions.ChannelHandler{channelEvents, ChannelManager}
	messageHandler := &actions.MessageHandler{messageEvents, ChannelManager}

	EventBroker.Resources[actions.Channel] = channelEvents
	EventBroker.Resources[actions.Message] = messageEvents

	go channelsHandler.Run()
	go messageHandler.Run()

}
