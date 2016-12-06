package main

import (
	"log"
	"net/http"
	ctrl "github.com/imdoroshenko/go-chat/controllers"
	m "github.com/imdoroshenko/go-chat/middleware"
	b "github.com/imdoroshenko/go-chat/bootstrap"
)

func main() {
	session := m.Session(b.Redis)
	composer := m.Adapt(&ctrl.Messenger{b.ChannelManager, b.EventBroker}, session)
	http.Handle("/", composer)

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
