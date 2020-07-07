package rest

import (
	"errors"
	"github.com/gorilla/schema"
	"github.com/julienschmidt/httprouter"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/composer"
	e "github.com/statistico/statistico-web-gateway/internal/app/errors"
	"net/http"
	"strconv"
)

type ResultHandler struct {
	composer composer.ResultComposer
}

func(h *ResultHandler) ByTeam(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	param := ps.ByName("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		failResponse(w, http.StatusBadRequest, err)
		return
	}

	filters := composer.Filters{}

	if err := schema.NewDecoder().Decode(&filters, r.URL.Query()); err != nil {
		failResponse(w, http.StatusUnprocessableEntity, errors.New("error parsing query parameters"))
		return
	}

	results, err := h.composer.ForTeam(uint64(id), &filters)

	if err != nil {
		if err == e.ErrorInternalServerError {
			errorResponse(w, http.StatusInternalServerError, errors.New("internal server error"))
			return
		}

		if err == e.ErrorBadGateway {
			errorResponse(w, http.StatusBadGateway, errors.New("bad gateway"))
			return
		}

		failResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	resultsResponse(w, results)
}

func resultsResponse(w http.ResponseWriter, results []*app.Result) {
	payload := struct {
		Results []*app.Result `json:"results"`
	}{}

	payload.Results = results

	successResponse(w, http.StatusOK, payload)
}

func NewResultHandler(c composer.ResultComposer) *ResultHandler {
	return &ResultHandler{composer: c}
}
