package main

import (
	"log"
	"net/http"
	ctrl "github.com/imdoroshenko/go-chat/controllers"
	m "github.com/imdoroshenko/go-chat/middleware"
)

func main() {
	http.Handle("/", &m.Session{&ctrl.Messenger{}})
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
