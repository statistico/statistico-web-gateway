package rest

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Server struct {
	router *httprouter.Router
	middleware []http.Handler
}

func (s *Server) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")
	}

	s.router.ServeHTTP(w, r)
}

func NewServer(router *httprouter.Router) *Server {
	return &Server{router: router}
}
