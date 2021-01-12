package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/statistico/statistico-grpc-gateway/internal/app/rest"
	"github.com/statistico/statistico-grpc-gateway/internal/bootstrap"
	"log"
	"net/http"
)

func main() {
	container := bootstrap.BuildContainer(bootstrap.BuildConfig())

	router := httprouter.New()
	
	router.GET("/", rest.RoutePath)
	router.GET("/healthcheck", rest.HealthCheck)
	router.GET("/openapi.json", rest.RenderApiDocs)

	router.GET("/competition/:id/seasons", container.RestSeasonHandler().ByCompetitionID)
	router.GET("/country/:id/competitions", container.RestCompetitionHandler().ByCountryID)
	router.POST("/result-search", container.RestResultHandler().Fetch)
	router.GET("/season/:id/teams", container.RestTeamHandler().BySeasonID)
	router.GET("/team/:id", container.RestTeamHandler().ByID)
	router.GET("/team/:id/seasons", container.RestSeasonHandler().ByTeamID)

	server := rest.MiddlewarePipe(
		router,
		rest.CorsMiddleware,
	)

	log.Fatal(http.ListenAndServe(":80", server))
}
