package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/statistico/statistico-web-gateway/internal/app/rest"
	"github.com/statistico/statistico-web-gateway/internal/bootstrap"
	"log"
	"net/http"
)

func main() {
	container := bootstrap.BuildContainer(bootstrap.BuildConfig())

	router := httprouter.New()
	
	router.GET("/", rest.RoutePath)
	router.GET("/healthcheck", rest.HealthCheck)
	router.GET("/openapi.json", rest.RenderApiDocs)
	
	router.GET("/team/:id", container.RestTeamHandler().TeamById)
	router.POST("/result-search", container.RestResultHandler().Fetch)

	log.Fatal(http.ListenAndServe(":80", &Server{router}))
}

type Server struct {
	r *httprouter.Router
}


func (s *Server) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")
	s.r.ServeHTTP(w, r)
}