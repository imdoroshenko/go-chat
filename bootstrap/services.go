package bootstrap

import (
	"github.com/imdoroshenko/go-chat/services"
)

const defaultChannel = "default"

var (
	ChannelManager = services.NewChannelManager()
	EventBroker = services.NewEventBroker()
)

func init() {
	EventBroker.ChannelManager = ChannelManager
	go EventBroker.Run()
	go ChannelManager.Run()
	ChannelManager.Open<- defaultChannel
	ChannelManager.Open<- "other"
}
