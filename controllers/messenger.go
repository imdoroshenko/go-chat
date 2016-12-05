package controllers

import (
	"net/http"
	"log"
	"github.com/gorilla/websocket"
	"github.com/imdoroshenko/go-chat/models"
	"github.com/imdoroshenko/go-chat/services"
)

const (
	defaultChannel = "default"
	socketBufferSize  = 1024
)

var (
	upgrader *websocket.Upgrader
	cm *services.ChannelManager
	eb *services.EventBroker
)

func init() {
	upgrader = &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
		ReadBufferSize: socketBufferSize,
		WriteBufferSize: socketBufferSize}

	cm = services.NewChannelManager()
	eb = services.NewMessageBroker()
	eb.ChannelManager = cm
	go eb.Run()
	go cm.Run()
	cm.Open<- defaultChannel
	cm.Open<- "other"
}

type Messenger struct {}

func (m *Messenger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
	}
	client := models.NewClient(conn,cm.Channels[defaultChannel], eb.Incoming)
	go client.Run()
	cm.Channels[defaultChannel].Join <- client
}
