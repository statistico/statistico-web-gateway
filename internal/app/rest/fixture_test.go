package rest_test

import (
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/composer"
	"github.com/statistico/statistico-web-gateway/internal/app/mock"
	"github.com/statistico/statistico-web-gateway/internal/app/rest"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestFixtureHandler_FixtureSearch(t *testing.T) {
	t.Run("returns 200 response containing fixture data", func(t *testing.T) {
		t.Helper()

		reader := strings.NewReader(`{
			"leagueIds": [12, 23], 
			"dateTo": "2020-06-11T13:00:00+00:00", 
			"dateFrom": "2020-06-11T13:00:00+00:00", 
			"statFilter": {
				"type": "goals", 
				"team": "home", 
				"metric": "total", 
				"measure": "gte", 
				"value": 2.75, 
				"venue": "home", 
				"games": 4
			}
		}`)

		req, err := http.NewRequest("POST", "/fixture-search", reader)

		if err != nil {
			t.Fatal(err)
		}

		s := new(mock.FixtureSearch)
		handler := rest.FixtureHandler{Composer: s}

		rr := httptest.NewRecorder()

		s.On("Search", fixtureQuery()).Return(fixtureSlice())

		handler.FixtureSearch(rr, req, nil)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status %v, Got status %v", http.StatusOK, rr.Code)
		}

		expected := "{\"status\":\"success\",\"data\":{\"fixtures\":[{\"id\":1234,\"competition\":{\"id\":567,\"name\":" +
			"\"English Premier League\",\"isCup\":false},\"season\":{\"id\":89,\"name\":\"2019/2020\",\"isCurrent\":true}," +
			"\"round\":{\"id\":2,\"name\":\"2\",\"seasonId\":89,\"startDate\":\"2020-06-11T13:00:00Z\",\"endDate\":\"2020-06-11T13:00:00Z\"}," +
			"\"homeTeam\":{\"id\":1,\"name\":\"West Ham United\"},\"awayTeam\":{\"id\":10,\"name\":\"Newcastle United\"}," +
			"\"venue\":{\"id\":4,\"name\":\"London Stadium\"},\"date\":\"2020-06-11T13:00:00Z\"}]}}"

		if rr.Body.String() != expected {
			t.Errorf("Expected body %s, got %s", expected, rr.Body.String())
		}
	})
}

func fixtureSlice() []*app.Fixture {
	fixtures := []*app.Fixture{}

	d, _ := time.Parse(time.RFC3339, "2020-06-11T13:00:00+00:00")

	fix := &app.Fixture{
		ID: 1234,
		Competition: app.Competition{
			ID:    567,
			Name:  "English Premier League",
			IsCup: false,
		},
		Season: app.Season{
			ID:        89,
			Name:      "2019/2020",
			IsCurrent: true,
		},
		Round: app.Round{
			ID:        2,
			Name:      "2",
			SeasonID:  89,
			StartDate: app.JsonDate(d),
			EndDate:   app.JsonDate(d),
		},
		HomeTeam: app.Team{
			ID:   1,
			Name: "West Ham United",
		},
		AwayTeam: app.Team{
			ID:   10,
			Name: "Newcastle United",
		},
		Venue: app.Venue{
			ID:   4,
			Name: "London Stadium",
		},
		Date: app.JsonDate(d),
	}

	fixtures = append(fixtures, fix)

	return fixtures
}

func fixtureQuery() *composer.FixtureSearchQuery {
	d, _ := time.Parse(time.RFC3339, "2020-06-11T13:00:00+00:00")

	return &composer.FixtureSearchQuery{
		LeagueIds: []uint64{12, 23},
		DateFrom:  d,
		DateTo:    d,
		FixtureStatFilter: composer.FixtureStatFilter{
			Type:    "goals",
			Team:    "home",
			Metric:  "total",
			Measure: "gte",
			Value:   2.75,
			Venue:   "home",
			Games:   4,
		},
	}
}
