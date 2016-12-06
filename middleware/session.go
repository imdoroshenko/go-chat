package middleware

import (
	"net/http"
	_ "github.com/mediocregopher/radix.v2/pool"
	"log"
)


func Session(str string) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("before", str)
			defer log.Println("after", str)
			h.ServeHTTP(w, r)
		})
	}
}
