package rest

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/statistico/statistico-grpc-gateway/internal/app"
	"github.com/statistico/statistico-grpc-gateway/internal/app/composer"
	"net/http"
)

type ResultHandler struct {
	composer composer.ResultComposer
}

func (h *ResultHandler) Fetch(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	filters := composer.ResultFilters{}

	if err := json.NewDecoder(r.Body).Decode(&filters); err != nil {
		failResponse(
			w,
			http.StatusBadRequest,
			fmt.Errorf("error parsing request body: %s", err.Error()),
		)
		return
	}

	results, err := h.composer.FetchResults(&filters)

	if err != nil {
		handleError(w, err)
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
