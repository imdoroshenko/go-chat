package controllers

import (
	"net/http"
	"log"
	"github.com/gorilla/websocket"
	"github.com/imdoroshenko/go-chat/models"
	"github.com/imdoroshenko/go-chat/services"
)

const defaultChannel = "default"

type Channel struct {
	ChannelManager *services.ChannelManager
	MessageBroker *services.MessageBroker
}

const socketBufferSize  = 1024

var upgrader = &websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
	ReadBufferSize:socketBufferSize,
	WriteBufferSize: socketBufferSize}

func (c *Channel) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
	}
	client := models.NewClient(conn, c.ChannelManager.Channels[defaultChannel], c.MessageBroker.Incoming)
	go client.Run()
	c.ChannelManager.Channels[defaultChannel].Join <- client
}
