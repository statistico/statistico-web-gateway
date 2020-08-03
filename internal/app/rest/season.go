package rest

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/composer"
	e "github.com/statistico/statistico-web-gateway/internal/app/errors"
	"net/http"
	"strconv"
)

type SeasonHandler struct {
	competitionComposer composer.CompetitionComposer
	seasonComposer composer.SeasonComposer
}

func (s *SeasonHandler) ByCompetitionId(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	param := ps.ByName("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		failResponse(w, http.StatusBadRequest, errors.New("unable to parse ID parameter as correct schema"))
		return
	}

	sort := r.URL.Query().Get("sort")

	seasons, err := s.competitionComposer.CompetitionSeasons(uint64(id), sort)

	if err != nil {
		if err == e.ErrorInternalServerError {
			errorResponse(w, http.StatusInternalServerError, errors.New("internal server error"))
			return
		}

		errorResponse(w, http.StatusBadGateway, errors.New("bad gateway"))
		return
	}

	seasonsResponse(w, seasons)
}

func (s *SeasonHandler) ByTeamId(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	param := ps.ByName("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		failResponse(w, http.StatusBadRequest, errors.New("unable to parse ID parameter as correct schema"))
		return
	}

	sort := r.URL.Query().Get("sort")

	seasons, err := s.seasonComposer.ByTeamId(uint64(id), sort)

	if err != nil {
		if err == e.ErrorInternalServerError {
			errorResponse(w, http.StatusInternalServerError, errors.New("internal server error"))
			return
		}

		errorResponse(w, http.StatusBadGateway, errors.New("bad gateway"))
		return
	}

	seasonsResponse(w, seasons)
}

func seasonsResponse(w http.ResponseWriter, seasons []*app.Season) {
	payload := struct {
		Seasons []*app.Season `json:"seasons"`
	}{}

	payload.Seasons = seasons

	successResponse(w, http.StatusOK, payload)
}

func NewSeasonHandler(c composer.CompetitionComposer, s composer.SeasonComposer) *SeasonHandler {
	return &SeasonHandler{competitionComposer: c, seasonComposer: s}
}
