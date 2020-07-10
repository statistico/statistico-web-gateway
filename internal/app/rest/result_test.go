package rest_test

import (
	"bytes"
	"errors"
	"github.com/julienschmidt/httprouter"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/composer"
	e "github.com/statistico/statistico-web-gateway/internal/app/errors"
	"github.com/statistico/statistico-web-gateway/internal/app/mock"
	"github.com/statistico/statistico-web-gateway/internal/app/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestResultHandler_ByTeam(t *testing.T) {
	t.Run("returns a 200 response containing results data", func(t *testing.T) {
		t.Helper()

		c := new(mock.ResultComposer)
		handler := rest.NewResultHandler(c)

		body := []byte(`{"dateBefore":"2020-01-01T00:00:00+00:00","dateAfter":"2020-01-01T12:00:00+00:00","limit":10,
			"sort":"date_desc","team":{"id":33,"venue":"home"},"seasonIds":[16036,12968]}`)
		req := httptest.NewRequest(http.MethodPost, "/result-search", bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		seasonIds := []uint64{16036, 12968}
		limit := uint64(10)
		sort := "date_desc"
		venue := "home"
		before, _ := time.Parse(time.RFC3339, "2020-01-01T00:00:00+00:00")
		after, _ := time.Parse(time.RFC3339, "2020-01-01T12:00:00+00:00")

		filters := composer.Filters{
			Limit:      &limit,
			DateBefore: &before,
			DateAfter:  &after,
			SeasonIds:  &seasonIds,
			Sort:       &sort,
			Team: &composer.TeamFilter{
				ID:    33,
				Venue: &venue,
			},
		}

		c.On("FetchResults", &filters).Return(results(), nil)

		handler.Fetch(res, req, httprouter.Params{})

		expected := `{"status":"success","data":{"results":[{"id":78102,"homeTeam":{"id":1,"name":"West Ham United","shortCode":null,` +
			`"countryId":8,"venueId":214,"nationalTeam":false,"founded":null,"logo":null},"awayTeam":{"id":10,"name":"Nottingham Forest",` +
			`"shortCode":null,"countryId":8,"venueId":300,"nationalTeam":false,"founded":null,"logo":null},"season":{"id":16036,"name":"2019/2020",` +
			`"isCurrent":false},"round":{"id":38,"name":"38","seasonId":16036,"startDate":"2020-07-07T12:00:00Z","endDate":"2020-07-23T23:59:59Z"},` +
			`"venue":{"id":214,"name":"London Stadium"},"date":"2020-07-07T15:00:00Z","stats":{"homeScore":5,"awayScore":2}}]}}`

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})

	t.Run("returns a 200 response handling less request body parameters", func(t *testing.T) {
		t.Helper()

		c := new(mock.ResultComposer)
		handler := rest.NewResultHandler(c)

		body := []byte(`{"sort":"date_desc","team":{"id":1}}`)
		req := httptest.NewRequest(http.MethodPost, "/result-search", bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		sort := "date_desc"

		filters := composer.Filters{
			Sort: &sort,
			Team: &composer.TeamFilter{ID: 1},
		}

		c.On("FetchResults", &filters).Return(results(), nil)

		handler.Fetch(res, req, httprouter.Params{})

		expected := `{"status":"success","data":{"results":[{"id":78102,"homeTeam":{"id":1,"name":"West Ham United","shortCode":null,` +
			`"countryId":8,"venueId":214,"nationalTeam":false,"founded":null,"logo":null},"awayTeam":{"id":10,"name":"Nottingham Forest",` +
			`"shortCode":null,"countryId":8,"venueId":300,"nationalTeam":false,"founded":null,"logo":null},"season":{"id":16036,"name":"2019/2020",` +
			`"isCurrent":false},"round":{"id":38,"name":"38","seasonId":16036,"startDate":"2020-07-07T12:00:00Z","endDate":"2020-07-23T23:59:59Z"},` +
			`"venue":{"id":214,"name":"London Stadium"},"date":"2020-07-07T15:00:00Z","stats":{"homeScore":5,"awayScore":2}}]}}`

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})

	t.Run("returns a 400 response if error parsing request body", func(t *testing.T) {
		t.Helper()

		c := new(mock.ResultComposer)
		handler := rest.NewResultHandler(c)

		body := []byte(`{"sort":10,"team":{"id":1}`)
		req := httptest.NewRequest(http.MethodPost, "/result-search", bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		sort := "date_desc"

		filters := composer.Filters{
			Sort: &sort,
			Team: &composer.TeamFilter{ID: 1},
		}

		c.AssertNotCalled(t, "FetchResults", &filters)

		handler.Fetch(res, req, httprouter.Params{})

		expected := `{"status":"fail","data":[{"message":"error parsing request body: unexpected EOF","code":1}]}`

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})

	t.Run("return 500 response if internal server error returned by client", func(t *testing.T) {
		t.Helper()

		c := new(mock.ResultComposer)
		handler := rest.NewResultHandler(c)

		body := []byte(`{"sort":"date_desc","team":{"id":1}}`)
		req := httptest.NewRequest(http.MethodPost, "/result-search", bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		sort := "date_desc"

		filters := composer.Filters{
			Sort: &sort,
			Team: &composer.TeamFilter{ID: 1},
		}

		c.On("FetchResults", &filters).Return([]*app.Result{}, e.ErrorInternalServerError)

		handler.Fetch(res, req, httprouter.Params{})

		expected := `{"status":"error","data":[{"message":"internal server error","code":1}]}`

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})

	t.Run("return 502 response if bad gateway error returned by client", func(t *testing.T) {
		t.Helper()

		c := new(mock.ResultComposer)
		handler := rest.NewResultHandler(c)

		body := []byte(`{"sort":"date_desc","team":{"id":1}}`)
		req := httptest.NewRequest(http.MethodPost, "/result-search", bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		sort := "date_desc"

		filters := composer.Filters{
			Sort: &sort,
			Team: &composer.TeamFilter{ID: 1},
		}

		c.On("FetchResults", &filters).Return([]*app.Result{}, e.ErrorBadGateway)

		handler.Fetch(res, req, httprouter.Params{})

		expected := `{"status":"error","data":[{"message":"bad gateway","code":1}]}`

		assert.Equal(t, http.StatusBadGateway, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})

	t.Run("returns 422 response if validation error returned by client", func(t *testing.T) {
		t.Helper()

		c := new(mock.ResultComposer)
		handler := rest.NewResultHandler(c)

		body := []byte(`{"sort":"date_desc","team":{"id":1}}`)
		req := httptest.NewRequest(http.MethodPost, "/result-search", bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		sort := "date_desc"

		filters := composer.Filters{
			Sort: &sort,
			Team: &composer.TeamFilter{ID: 1},
		}

		c.On("FetchResults", &filters).Return([]*app.Result{}, errors.New("validation error"))

		handler.Fetch(res, req, httprouter.Params{})

		expected := `{"status":"fail","data":[{"message":"validation error","code":1}]}`

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})
}

func results() []*app.Result {
	var results []*app.Result

	home := app.Team{
		ID:        1,
		Name:      "West Ham United",
		CountryID: 8,
		VenueID:   214,
	}

	away := app.Team{
		ID:        10,
		Name:      "Nottingham Forest",
		CountryID: 8,
		VenueID:   300,
	}

	season := app.Season{
		ID:   16036,
		Name: "2019/2020",
	}

	start, _ := time.Parse(time.RFC3339, "2020-07-07T12:00:00+00:00")
	end, _ := time.Parse(time.RFC3339, "2020-07-23T23:59:59+00:00")

	round := app.Round{
		ID:        38,
		Name:      "38",
		SeasonID:  16036,
		StartDate: app.JsonDate(start),
		EndDate:   app.JsonDate(end),
	}

	venue := app.Venue{
		ID:   214,
		Name: "London Stadium",
	}

	stats := app.ResultStats{
		HomeScore: 5,
		AwayScore: 2,
	}

	date, _ := time.Parse(time.RFC3339, "2020-07-07T15:00:00+00:00")

	result := app.Result{
		ID:       78102,
		HomeTeam: home,
		AwayTeam: away,
		Season:   season,
		Round:    round,
		Venue:    venue,
		DateTime: app.JsonDate(date),
		Stats:    stats,
	}

	results = append(results, &result)

	return results
}
