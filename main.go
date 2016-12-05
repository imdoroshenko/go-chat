package main

import (
	"log"
	"net/http"
	"github.com/imdoroshenko/go-chat/controllers"
)

func main() {
	http.Handle("/", &controllers.Messenger{})
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
