package rest

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"github.com/statistico/statistico-grpc-gateway/internal/app"
	"github.com/statistico/statistico-grpc-gateway/internal/app/composer"
	e "github.com/statistico/statistico-grpc-gateway/internal/app/errors"
	"net/http"
	"strconv"
)

type CompetitionHandler struct {
	composer composer.CompetitionComposer
}

func (c *CompetitionHandler) ByCountryID(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	param := ps.ByName("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		failResponse(w, http.StatusBadRequest, errors.New("unable to parse ID parameter as correct schema"))
		return
	}

	competitions, err := c.composer.CompetitionsByCountryID(uint64(id))

	if err != nil {
		if err == e.ErrorInternalServerError {
			errorResponse(w, http.StatusInternalServerError, errors.New("internal server error"))
			return
		}

		errorResponse(w, http.StatusBadGateway, errors.New("bad gateway"))
		return
	}

	competitionsResponse(w, competitions)
}

func competitionsResponse(w http.ResponseWriter, competitions []*app.Competition) {
	payload := struct {
		Competitions []*app.Competition `json:"competitions"`
	}{}

	payload.Competitions = competitions

	successResponse(w, http.StatusOK, payload)
}

func NewCompetitionHandler(c composer.CompetitionComposer) *CompetitionHandler {
	return &CompetitionHandler{composer: c}
}
