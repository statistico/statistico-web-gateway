package rest

import (
	"net/http"
)

func MiddlewarePipe(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, m := range middleware {
		h = m(h)
	}

	return h
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r*http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "*")
			return
		}

		next.ServeHTTP(w, r)
	})
}
