package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/composer"
	e "github.com/statistico/statistico-web-gateway/internal/app/errors"
	"net/http"
)

type ResultHandler struct {
	composer composer.ResultComposer
}

func (h *ResultHandler) Fetch(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	filters := composer.Filters{}

	if err := json.NewDecoder(r.Body).Decode(&filters); err != nil {
		failResponse(
			w,
			http.StatusUnprocessableEntity,
			fmt.Errorf("error parsing request body: %s", err.Error()),
		)
		return
	}

	results, err := h.composer.FetchResults(&filters)

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
