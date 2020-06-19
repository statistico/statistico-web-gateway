package rest

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/composer"
	"net/http"
)

type FixtureHandler struct {
	Composer composer.FixtureHandler
}

func (f FixtureHandler) FixtureSearch(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var q *composer.FixtureSearchQuery

	err := json.NewDecoder(r.Body).Decode(&q)

	if err != nil {
		code, err := parseJsonError(err)
		failResponse(w, code, err)
		return
	}

	fixtures, err := f.Composer.Search(q)

	if err != nil {
		failResponse(w, 500, err)
		return
	}

	res := fixtureResponse{
		Fixtures: fixtures,
	}

	successResponse(w, http.StatusOK, res)
}

type fixtureResponse struct {
	Fixtures map[uint64][]*app.Fixture `json:"fixtures"`
}
