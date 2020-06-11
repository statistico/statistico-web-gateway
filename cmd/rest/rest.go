package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/statistico/statistico-web-gateway/internal/app/composer"
	"github.com/statistico/statistico-web-gateway/internal/app/rest"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", rest.RoutePath)
	router.GET("/healthcheck", rest.HealthCheck)
	router.GET("/openapi.json", rest.RenderApiDocs)

	h := rest.FixtureHandler{Composer: &composer.FixtureSearch{}}

	router.POST("/fixture-search", h.FixtureSearch)

	log.Fatal(http.ListenAndServe(":80", router))
}