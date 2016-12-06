package middleware

import (
	"net/http"
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
)


func Session(pool pool.Pool) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		})
	}
}

func checkSession(redis redis.Client) {

}

func setSession(redis redis.Client) {

}
