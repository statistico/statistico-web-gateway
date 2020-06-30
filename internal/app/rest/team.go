package rest

import (
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
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

	successResponse(w, http.StatusOK, team)
}

func NewTeamHandler(c composer.TeamComposer) *TeamHandler {
	return &TeamHandler{c}
}

func notFoundResponse(w http.ResponseWriter, id string) {
	failResponse(w, http.StatusNotFound, errors.New(fmt.Sprintf("team with id '%s' does not exist", id)))
}
