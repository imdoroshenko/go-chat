package main

import (
	"log"
	"net/http"
)

const defaultChannel = "default"

var (
	cm *ChannelManager = NewChannelManager()
	mb *MessageBroker = NewMessageBroker()
)

type mainCtrl struct {}

func (m *mainCtrl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
	}
	client := newClient(conn, cm.Channels["default"], mb.Incoming)
	go client.Run()
	cm.Channels[defaultChannel].Join <- client
}

func main() {
	go mb.Run()
	go cm.Run()
	cm.Open<- defaultChannel
	cm.Open<- "other"
	http.Handle("/", &mainCtrl{})
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
