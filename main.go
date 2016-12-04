package main

import (
	"log"
	"net/http"
)


//type mainCtrl struct {}
//
//func (m *mainCtrl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	log.Println("Main ctrl")
//	conn, err := upgrader.Upgrade(w, r, nil)
//	if err != nil {
//		log.Fatal("ServeHTTP:", err)
//		return
//	}
//
//	go func () {
//		for {
//			messageType, p, err := conn.ReadMessage()
//			if err != nil {
//				return
//			}
//			if err = conn.WriteMessage(messageType, p); err != nil {
//				return
//			}
//		}
//	}()
//}

func main() {
	//http.Handle("/", &mainCtrl{})
	channel := newChannel("default")
	http.Handle("/", channel)
	go channel.run()
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
