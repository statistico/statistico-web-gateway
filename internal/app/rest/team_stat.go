package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/composer"
	"net/http"
)

type TeamStatHandler struct {
	composer composer.TeamStatComposer
}

func (h *TeamStatHandler) Fetch(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	filters := composer.TeamStatFilters{}

	if err := json.NewDecoder(r.Body).Decode(&filters); err != nil {
		failResponse(
			w,
			http.StatusBadRequest,
			fmt.Errorf("error parsing request body: %s", err.Error()),
		)
		return
	}

	stats, err := h.composer.FetchStats(context.Background(), &filters)

	if err != nil {
		handleError(w, err)
		return
	}

	teamStatResponse(w, stats)
}

func teamStatResponse(w http.ResponseWriter, stats []*app.TeamStat) {
	payload := struct {
		Stats []*app.TeamStat `json:"stats"`
	}{}

	payload.Stats = stats

	successResponse(w, http.StatusOK, payload)
}

func NewTeamStatHandler(c composer.TeamStatComposer) *TeamStatHandler {
	return &TeamStatHandler{composer: c}
}
