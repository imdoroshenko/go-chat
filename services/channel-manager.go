package services

import (
	"github.com/imdoroshenko/go-chat/models"
	"log"
)

type ChannelManager struct {
	Channels map[string]*models.Channel
	Open chan string
	Close chan string
}

func (cm *ChannelManager) Run() {
	for {
		select {
			case name := <-cm.Open:
				cm.Channels[name] = models.NewChannel(name)
				go cm.Channels[name].Run()
				log.Printf("%s channel is open\n", name)
			case name := <-cm.Close:
				cm.Channels[name].Stop <- 1
				delete(cm.Channels, name)
				log.Printf("%s channel is closed\n", name)
		}
	}
}

//
//func (cm *ChannelsManager) TransferClients(from string, to string) err {
//	if fromCh, ok := cm.Channels[from]; ok {
//		return fmt.Errorf("%s channel does not exist", from)
//	}
//	if toCh, ok := cm.Channels[from]; ok {
//		return fmt.Errorf("%s channel does not exist", to)
//	}
//	for client, _ := range fromCh.clients {
//		fromCh.le
//	}
//}

func NewChannelManager() *ChannelManager {
	return &ChannelManager{
		Channels: make(map[string]*models.Channel),
		Open:   make(chan string),
		Close:   make(chan string),
	}
}
