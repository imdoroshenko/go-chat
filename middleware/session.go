package middleware

import (
	"net/http"
	"log"
)

type Session struct {
	Next http.Handler
}

func (s *Session) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("before request")
	defer log.Println("after request")
	s.Next.ServeHTTP(w, r)
}

