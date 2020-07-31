package rest

import (
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/composer"
	e "github.com/statistico/statistico-web-gateway/internal/app/errors"
	"net/http"
	"strconv"
)

type TeamHandler struct {
	composer composer.TeamComposer
}

func (t *TeamHandler) TeamById(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	param := ps.ByName("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		notFoundResponse(w, param)
		return
	}

	team, err := t.composer.TeamById(uint64(id))

	if err != nil {
		if err == e.ErrorNotFound {
			notFoundResponse(w, param)
			return
		}

		errorResponse(w, http.StatusBadGateway, errors.New("bad gateway"))
		return
	}

	teamResponse(w, team)
}

func (t *TeamHandler) TeamsBySeasonId(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	param := ps.ByName("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		failResponse(w, http.StatusBadRequest, errors.New("unable to parse ID parameter as correct schema"))
		return
	}

	teams, err := t.composer.TeamsBySeasonId(uint64(id))

	if err != nil {
		if err == e.ErrorInternalServerError {
			errorResponse(w, http.StatusInternalServerError, errors.New("internal server error"))
			return
		}

		errorResponse(w, http.StatusBadGateway, errors.New("bad gateway"))
		return
	}

	teamsResponse(w, teams)
}

func NewTeamHandler(c composer.TeamComposer) *TeamHandler {
	return &TeamHandler{c}
}

func notFoundResponse(w http.ResponseWriter, id string) {
	failResponse(w, http.StatusNotFound, errors.New(fmt.Sprintf("team with id '%s' does not exist", id)))
}

func teamResponse(w http.ResponseWriter, team *app.Team) {
	payload := struct {
		Team *app.Team `json:"team"`
	}{}

	payload.Team = team

	successResponse(w, http.StatusOK, payload)
}

func teamsResponse(w http.ResponseWriter, teams []*app.Team) {
	payload := struct {
		Teams []*app.Team `json:"teams"`
	}{}

	payload.Teams = teams

	successResponse(w, http.StatusOK, payload)
}
