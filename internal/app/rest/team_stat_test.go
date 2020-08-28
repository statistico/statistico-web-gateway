package rest_test

import (
	"bytes"
	"context"
	"errors"
	"github.com/julienschmidt/httprouter"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/composer"
	e "github.com/statistico/statistico-web-gateway/internal/app/errors"
	"github.com/statistico/statistico-web-gateway/internal/app/mock"
	"github.com/statistico/statistico-web-gateway/internal/app/rest"
	"github.com/stretchr/testify/assert"
	mock2 "github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTeamStatHandler_Fetch(t *testing.T) {
	t.Run("returns a 200 response containing team stat data", func(t *testing.T) {
		t.Helper()

		c := new(mock.TeamStatComposer)
		handler := rest.NewTeamStatHandler(c)

		body := []byte(`{"dateBefore":"2020-01-01T00:00:00+00:00","dateAfter":"2020-01-01T12:00:00+00:00","limit":10,
			"sort":"date_desc","opponent":true,"team":{"id":33,"venue":"home"},"seasonIds":[16036,12968],"stat":"goals"}`)
		req := httptest.NewRequest(http.MethodPost, "/team-stat-search", bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		seasonIds := []uint64{16036, 12968}
		limit := uint64(10)
		sort := "date_desc"
		venue := "home"
		before, _ := time.Parse(time.RFC3339, "2020-01-01T00:00:00+00:00")
		after, _ := time.Parse(time.RFC3339, "2020-01-01T12:00:00+00:00")

		args := mock2.MatchedBy(func (f *composer.TeamStatFilters) bool {
			a := assert.New(t)
			a.Equal(seasonIds, *f.SeasonIds)
			a.Equal(before, *f.DateBefore)
			a.Equal(after, *f.DateAfter)
			a.Equal(sort, *f.Sort)
			a.Equal(limit, *f.Limit)
			a.Equal(uint64(33), f.Team.ID)
			a.Equal(venue, *f.Team.Venue)
			a.Equal(true, *f.Opponent)
			return true
		})

		c.On("FetchStats", context.Background(), args).Return(teamStats(), nil)

		handler.Fetch(res, req, httprouter.Params{})

		expected := `{"status":"success","data":{"stats":[{"fixtureId":491,"stat":"goals","value":2}]}}`

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})

	t.Run("returns a 200 response containing team stat data handling null stat values", func(t *testing.T) {
		t.Helper()

		c := new(mock.TeamStatComposer)
		handler := rest.NewTeamStatHandler(c)

		body := []byte(`{"dateBefore":"2020-01-01T00:00:00+00:00","dateAfter":"2020-01-01T12:00:00+00:00","limit":10,
			"sort":"date_desc","opponent":true,"team":{"id":33,"venue":"home"},"seasonIds":[16036,12968],"stat":"goals"}`)
		req := httptest.NewRequest(http.MethodPost, "/team-stat-search", bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		seasonIds := []uint64{16036, 12968}
		limit := uint64(10)
		sort := "date_desc"
		venue := "home"
		before, _ := time.Parse(time.RFC3339, "2020-01-01T00:00:00+00:00")
		after, _ := time.Parse(time.RFC3339, "2020-01-01T12:00:00+00:00")

		args := mock2.MatchedBy(func (f *composer.TeamStatFilters) bool {
			a := assert.New(t)
			a.Equal(seasonIds, *f.SeasonIds)
			a.Equal(before, *f.DateBefore)
			a.Equal(after, *f.DateAfter)
			a.Equal(sort, *f.Sort)
			a.Equal(limit, *f.Limit)
			a.Equal(uint64(33), f.Team.ID)
			a.Equal(venue, *f.Team.Venue)
			a.Equal(true, *f.Opponent)
			return true
		})

		var stats []*app.TeamStat

		stat := app.TeamStat{
			FixtureID: 491,
			Stat:      "goals",
		}

		stats = append(stats, &stat)

		c.On("FetchStats", context.Background(), args).Return(stats, nil)

		handler.Fetch(res, req, httprouter.Params{})

		expected := `{"status":"success","data":{"stats":[{"fixtureId":491,"stat":"goals","value":null}]}}`

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})

	t.Run("returns a 400 response if error parsing request body", func(t *testing.T) {
		t.Helper()

		c := new(mock.TeamStatComposer)
		handler := rest.NewTeamStatHandler(c)

		body := []byte(`{"sort":10,"team":{"id":1}`)
		req := httptest.NewRequest(http.MethodPost, "/team-stat-search", bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		c.AssertNotCalled(t, "FetchStats")

		handler.Fetch(res, req, httprouter.Params{})

		expected := `{"status":"fail","data":[{"message":"error parsing request body: unexpected EOF","code":1}]}`

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})

	t.Run("return 500 response if internal server error returned by client", func(t *testing.T) {
		t.Helper()

		c := new(mock.TeamStatComposer)
		handler := rest.NewTeamStatHandler(c)

		body := []byte(`{"dateBefore":"2020-01-01T00:00:00+00:00","dateAfter":"2020-01-01T12:00:00+00:00","limit":10,
			"sort":"date_desc","opponent":true,"team":{"id":33,"venue":"home"},"seasonIds":[16036,12968],"stat":"goals"}`)
		req := httptest.NewRequest(http.MethodPost, "/team-stat-search", bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		seasonIds := []uint64{16036, 12968}
		limit := uint64(10)
		sort := "date_desc"
		venue := "home"
		before, _ := time.Parse(time.RFC3339, "2020-01-01T00:00:00+00:00")
		after, _ := time.Parse(time.RFC3339, "2020-01-01T12:00:00+00:00")

		args := mock2.MatchedBy(func (f *composer.TeamStatFilters) bool {
			a := assert.New(t)
			a.Equal(seasonIds, *f.SeasonIds)
			a.Equal(before, *f.DateBefore)
			a.Equal(after, *f.DateAfter)
			a.Equal(sort, *f.Sort)
			a.Equal(limit, *f.Limit)
			a.Equal(uint64(33), f.Team.ID)
			a.Equal(venue, *f.Team.Venue)
			a.Equal(true, *f.Opponent)
			return true
		})

		c.On("FetchStats", context.Background(), args).Return([]*app.TeamStat{}, e.ErrorInternalServerError)

		handler.Fetch(res, req, httprouter.Params{})

		expected := `{"status":"error","data":[{"message":"internal server error","code":1}]}`

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})

	t.Run("return 502 response if bad gateway error returned by client", func(t *testing.T) {
		t.Helper()

		c := new(mock.TeamStatComposer)
		handler := rest.NewTeamStatHandler(c)

		body := []byte(`{"dateBefore":"2020-01-01T00:00:00+00:00","dateAfter":"2020-01-01T12:00:00+00:00","limit":10,
			"sort":"date_desc","opponent":true,"team":{"id":33,"venue":"home"},"seasonIds":[16036,12968],"stat":"goals"}`)
		req := httptest.NewRequest(http.MethodPost, "/team-stat-search", bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		seasonIds := []uint64{16036, 12968}
		limit := uint64(10)
		sort := "date_desc"
		venue := "home"
		before, _ := time.Parse(time.RFC3339, "2020-01-01T00:00:00+00:00")
		after, _ := time.Parse(time.RFC3339, "2020-01-01T12:00:00+00:00")

		args := mock2.MatchedBy(func (f *composer.TeamStatFilters) bool {
			a := assert.New(t)
			a.Equal(seasonIds, *f.SeasonIds)
			a.Equal(before, *f.DateBefore)
			a.Equal(after, *f.DateAfter)
			a.Equal(sort, *f.Sort)
			a.Equal(limit, *f.Limit)
			a.Equal(uint64(33), f.Team.ID)
			a.Equal(venue, *f.Team.Venue)
			a.Equal(true, *f.Opponent)
			return true
		})

		c.On("FetchStats", context.Background(), args).Return([]*app.TeamStat{}, e.ErrorBadGateway)

		handler.Fetch(res, req, httprouter.Params{})

		expected := `{"status":"error","data":[{"message":"bad gateway","code":1}]}`

		assert.Equal(t, http.StatusBadGateway, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})

	t.Run("returns 422 response if validation error returned by client", func(t *testing.T) {
		t.Helper()

		c := new(mock.TeamStatComposer)
		handler := rest.NewTeamStatHandler(c)

		body := []byte(`{"dateBefore":"2020-01-01T00:00:00+00:00","dateAfter":"2020-01-01T12:00:00+00:00","limit":10,
			"sort":"date_desc","opponent":true,"team":{"id":33,"venue":"home"},"seasonIds":[16036,12968],"stat":"goals"}`)
		req := httptest.NewRequest(http.MethodPost, "/team-stat-search", bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		seasonIds := []uint64{16036, 12968}
		limit := uint64(10)
		sort := "date_desc"
		venue := "home"
		before, _ := time.Parse(time.RFC3339, "2020-01-01T00:00:00+00:00")
		after, _ := time.Parse(time.RFC3339, "2020-01-01T12:00:00+00:00")

		args := mock2.MatchedBy(func (f *composer.TeamStatFilters) bool {
			a := assert.New(t)
			a.Equal(seasonIds, *f.SeasonIds)
			a.Equal(before, *f.DateBefore)
			a.Equal(after, *f.DateAfter)
			a.Equal(sort, *f.Sort)
			a.Equal(limit, *f.Limit)
			a.Equal(uint64(33), f.Team.ID)
			a.Equal(venue, *f.Team.Venue)
			a.Equal(true, *f.Opponent)
			return true
		})

		c.On("FetchStats", context.Background(), args).Return([]*app.TeamStat{}, errors.New("validation error"))

		handler.Fetch(res, req, httprouter.Params{})

		expected := `{"status":"fail","data":[{"message":"validation error","code":1}]}`

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})
}

func teamStats() []*app.TeamStat {
	var stats []*app.TeamStat

	goals := uint32(2)

	stat := app.TeamStat{
		FixtureID: 491,
		Stat:      "goals",
		Value:     &goals,
	}

	stats = append(stats, &stat)

	return stats
}
