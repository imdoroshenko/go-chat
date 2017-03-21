package controllers

import (
	"net/http"
	"log"
	"github.com/gorilla/websocket"
	"github.com/imdoroshenko/go-chat/models"
	"github.com/imdoroshenko/go-chat/services"
)

const socketBufferSize  = 1024

var upgrader = &websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
	ReadBufferSize: socketBufferSize,
	WriteBufferSize: socketBufferSize}

const defaultChannel = "default"

type Messenger struct {
	Cm *services.ChannelManager
	Eb *services.EventBroker
}

func (m *Messenger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("ServeHTTP: ", err)
		return
	}
	log.Println("New client connected")
	client := models.NewClient(conn, m.Cm.Channels[defaultChannel], m.Eb.Incoming)
	go client.Run()
	m.Cm.Channels[defaultChannel].Join <- client
}
