package rest_test

import (
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

		req := httptest.NewRequest(http.MethodGet, "/team/32/results", nil)
		q := req.URL.Query()
		q.Add("dateBefore", "2020-01-01T00:00:00+00:00")
		q.Add("dateAfter", "2020-01-01T12:00:00+00:00")
		q.Add("limit", "10")
		q.Add("sort", "date_desc")
		q.Add("venue", "home")
		req.URL.RawQuery = q.Encode()

		res := httptest.NewRecorder()
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "33",
			},
		}

		limit := uint64(10)
		sort := "date_desc"
		venue := "home"
		before, _ := time.Parse(time.RFC3339, "2020-01-01T00:00:00+00:00")
		after, _ := time.Parse(time.RFC3339, "2020-01-01T12:00:00+00:00")

		filters := composer.Filters{
			Limit:      &limit,
			DateBefore: &before,
			DateAfter:  &after,
			Sort:       &sort,
			Venue:      &venue,
		}

		c.On("ForTeam", uint64(33), &filters).Return(results(), nil)

		handler.ByTeam(res, req, params)

		expected := `{"status":"success","data":{"results":[{"id":78102,"homeTeam":{"id":1,"name":"West Ham United","shortCode":null,` +
			`"countryId":8,"venueId":214,"nationalTeam":false,"founded":null,"logo":null},"awayTeam":{"id":10,"name":"Nottingham Forest",` +
			`"shortCode":null,"countryId":8,"venueId":300,"nationalTeam":false,"founded":null,"logo":null},"season":{"id":16036,"name":"2019/2020",` +
			`"isCurrent":false},"round":{"id":38,"name":"38","seasonId":16036,"startDate":"2020-07-07T12:00:00Z","endDate":"2020-07-23T23:59:59Z"},` +
			`"venue":{"id":214,"name":"London Stadium"},"date":"2020-07-07T15:00:00Z","stats":{"homeScore":5,"awayScore":2}}]}}`

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})

	t.Run("returns a 200 response handling less query parameters", func(t *testing.T) {
		t.Helper()

		c := new(mock.ResultComposer)
		handler := rest.NewResultHandler(c)

		req := httptest.NewRequest(http.MethodGet, "/team/32/results", nil)
		q := req.URL.Query()
		q.Add("sort", "date_desc")
		q.Add("venue", "home")
		req.URL.RawQuery = q.Encode()

		res := httptest.NewRecorder()
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "33",
			},
		}

		sort := "date_desc"
		venue := "home"

		filters := composer.Filters{
			Sort:       &sort,
			Venue:      &venue,
		}

		c.On("ForTeam", uint64(33), &filters).Return(results(), nil)

		handler.ByTeam(res, req, params)

		expected := `{"status":"success","data":{"results":[{"id":78102,"homeTeam":{"id":1,"name":"West Ham United","shortCode":null,` +
			`"countryId":8,"venueId":214,"nationalTeam":false,"founded":null,"logo":null},"awayTeam":{"id":10,"name":"Nottingham Forest",` +
			`"shortCode":null,"countryId":8,"venueId":300,"nationalTeam":false,"founded":null,"logo":null},"season":{"id":16036,"name":"2019/2020",` +
			`"isCurrent":false},"round":{"id":38,"name":"38","seasonId":16036,"startDate":"2020-07-07T12:00:00Z","endDate":"2020-07-23T23:59:59Z"},` +
			`"venue":{"id":214,"name":"London Stadium"},"date":"2020-07-07T15:00:00Z","stats":{"homeScore":5,"awayScore":2}}]}}`

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})

	t.Run("returns a 422 response if error parsing query parameter schema", func(t *testing.T) {
		t.Helper()

		c := new(mock.ResultComposer)
		handler := rest.NewResultHandler(c)

		req := httptest.NewRequest(http.MethodGet, "/team/32/results", nil)
		q := req.URL.Query()
		q.Add("sorted", "date_desc")
		req.URL.RawQuery = q.Encode()

		res := httptest.NewRecorder()
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "33",
			},
		}

		c.AssertNotCalled(t,"ForTeam", uint64(33), &composer.Filters{})

		handler.ByTeam(res, req, params)

		expected := `{"status":"fail","data":[{"message":"error parsing query parameters","code":1}]}`

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})

	t.Run("return 500 response if internal server error returned by client", func(t *testing.T) {
		t.Helper()

		c := new(mock.ResultComposer)
		handler := rest.NewResultHandler(c)

		req := httptest.NewRequest(http.MethodGet, "/team/32/results", nil)
		q := req.URL.Query()
		q.Add("sort", "date_desc")
		q.Add("venue", "home")
		req.URL.RawQuery = q.Encode()

		res := httptest.NewRecorder()
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "33",
			},
		}

		sort := "date_desc"
		venue := "home"

		filters := composer.Filters{
			Sort:       &sort,
			Venue:      &venue,
		}

		c.On("ForTeam", uint64(33), &filters).Return([]*app.Result{}, e.ErrorInternalServerError)

		handler.ByTeam(res, req, params)

		expected := `{"status":"error","data":[{"message":"internal server error","code":1}]}`

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})

	t.Run("return 502 response if bad gateway error returned by client", func(t *testing.T) {
		t.Helper()

		c := new(mock.ResultComposer)
		handler := rest.NewResultHandler(c)

		req := httptest.NewRequest(http.MethodGet, "/team/32/results", nil)
		q := req.URL.Query()
		q.Add("sort", "date_desc")
		q.Add("venue", "home")
		req.URL.RawQuery = q.Encode()

		res := httptest.NewRecorder()
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "33",
			},
		}

		sort := "date_desc"
		venue := "home"

		filters := composer.Filters{
			Sort:       &sort,
			Venue:      &venue,
		}

		c.On("ForTeam", uint64(33), &filters).Return([]*app.Result{}, e.ErrorBadGateway)

		handler.ByTeam(res, req, params)

		expected := `{"status":"error","data":[{"message":"bad gateway","code":1}]}`

		assert.Equal(t, http.StatusBadGateway, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})

	t.Run("returns 422 response if validation error returned by client", func(t *testing.T) {
		t.Helper()

		c := new(mock.ResultComposer)
		handler := rest.NewResultHandler(c)

		req := httptest.NewRequest(http.MethodGet, "/team/32/results", nil)
		q := req.URL.Query()
		q.Add("sort", "date_desc")
		q.Add("venue", "home")
		req.URL.RawQuery = q.Encode()

		res := httptest.NewRecorder()
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "33",
			},
		}

		sort := "date_desc"
		venue := "home"

		filters := composer.Filters{
			Sort:       &sort,
			Venue:      &venue,
		}

		c.On("ForTeam", uint64(33), &filters).Return([]*app.Result{}, errors.New("validation error"))

		handler.ByTeam(res, req, params)

		expected := `{"status":"fail","data":[{"message":"validation error","code":1}]}`

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})
}

func results() []*app.Result {
	var results []*app.Result

	home := app.Team{
		ID:             1,
		Name:           "West Ham United",
		CountryID:      8,
		VenueID:        214,
	}

	away := app.Team{
		ID:             10,
		Name:           "Nottingham Forest",
		CountryID:      8,
		VenueID:        300,
	}

	season := app.Season{
		ID:        16036,
		Name:      "2019/2020",
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
		HomeScore:    5,
		AwayScore:    2,
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
