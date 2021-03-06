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

	router.GET("/competition/:id/seasons", container.RestSeasonHandler().ByCompetitionId)
	router.GET("/country/:id/competitions", container.RestCompetitionHandler().ByCountryId)
	router.POST("/result-search", container.RestResultHandler().Fetch)
	router.GET("/season/:id/teams", container.RestTeamHandler().TeamsBySeasonId)
	router.GET("/team/:id", container.RestTeamHandler().TeamById)
	router.GET("/team/:id/seasons", container.RestSeasonHandler().ByTeamId)
	router.POST("/team-stat-search", container.RestTeamStatHandler().Fetch)

	server := rest.MiddlewarePipe(
		router,
		rest.CorsMiddleware,
	)

	log.Fatal(http.ListenAndServe(":80", server))
}
