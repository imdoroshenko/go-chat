package main

import (
	"log"
	"net/http"
	ctrl "github.com/imdoroshenko/go-chat/controllers"
	"os"
	_ "net/http/pprof"
)

func main() {
	log.Println(os.Getpid())
	//session := m.Session(b.Redis)
	//composer := m.Adapt(&ctrl.Messenger{b.ChannelManager, b.EventBroker}, session)
	http.Handle("/", &ctrl.Messenger{ChannelManager, EventBroker})

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
