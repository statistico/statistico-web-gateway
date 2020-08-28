package rest_test

import (
	"github.com/julienschmidt/httprouter"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/errors"
	"github.com/statistico/statistico-web-gateway/internal/app/mock"
	"github.com/statistico/statistico-web-gateway/internal/app/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTeamHandler_TeamById(t *testing.T) {
	t.Run("returns a 200 response containing team data", func(t *testing.T) {
		t.Helper()

		composer := new(mock.TeamComposer)
		handler := rest.NewTeamHandler(composer)

		req := httptest.NewRequest(http.MethodGet, "/team/241", nil)
		res := httptest.NewRecorder()
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "241",
			},
		}

		short := "RFC"
		founded := uint64(1920)

		team := app.Team{
			ID:           241,
			Name:         "Romford FC",
			ShortCode:    &short,
			CountryID:    8,
			VenueID:      9182,
			NationalTeam: false,
			Founded:      &founded,
			Logo:         nil,
		}

		composer.On("TeamById", uint64(241)).Return(&team, nil)

		handler.TeamById(res, req, params)

		expected := `{"status":"success","data":{"team":{"id":241,"name":"Romford FC","shortCode":"RFC","countryId":8,"venueId":9182,` +
			`"nationalTeam":false,"founded":1920,"logo":null}}}`

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expected, res.Body.String())
		composer.AssertExpectations(t)
	})

	t.Run("returns 404 response if unable to parse id", func(t *testing.T) {
		t.Helper()

		composer := new(mock.TeamComposer)
		handler := rest.NewTeamHandler(composer)

		req := httptest.NewRequest(http.MethodGet, "/team/241", nil)
		res := httptest.NewRecorder()
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "Hello",
			},
		}

		composer.AssertNotCalled(t, "TeamById", uint64(241))

		handler.TeamById(res, req, params)

		expected := `{"status":"fail","data":[{"message":"team with id 'Hello' does not exist","code":1}]}`

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, expected, res.Body.String())
		composer.AssertExpectations(t)
	})

	t.Run("returns 404 response if not found response returned by composer", func(t *testing.T) {
		t.Helper()

		composer := new(mock.TeamComposer)
		handler := rest.NewTeamHandler(composer)

		req := httptest.NewRequest(http.MethodGet, "/team/241", nil)
		res := httptest.NewRecorder()
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "241",
			},
		}

		composer.On("TeamById", uint64(241)).Return(&app.Team{}, errors.ErrorNotFound)

		handler.TeamById(res, req, params)

		expected := `{"status":"fail","data":[{"message":"team with id '241' does not exist","code":1}]}`

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, expected, res.Body.String())
		composer.AssertExpectations(t)
	})

	t.Run("returns 502 response if bad gateway response returned by composer", func(t *testing.T) {
		t.Helper()

		composer := new(mock.TeamComposer)
		handler := rest.NewTeamHandler(composer)

		req := httptest.NewRequest(http.MethodGet, "/team/241", nil)
		res := httptest.NewRecorder()
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "241",
			},
		}

		composer.On("TeamById", uint64(241)).Return(&app.Team{}, errors.ErrorBadGateway)

		handler.TeamById(res, req, params)

		expected := `{"status":"error","data":[{"message":"bad gateway","code":1}]}`

		assert.Equal(t, http.StatusBadGateway, res.Code)
		assert.Equal(t, expected, res.Body.String())
		composer.AssertExpectations(t)
	})
}

func TestTeamHandler_TeamsBySeasonId(t *testing.T) {
	t.Run("returns a 200 response containing teams data", func(t *testing.T) {
		t.Helper()

		composer := new(mock.TeamComposer)
		handler := rest.NewTeamHandler(composer)

		req := httptest.NewRequest(http.MethodGet, "/season/16036/teams", nil)
		res := httptest.NewRecorder()
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "16036",
			},
		}

		short := "RFC"
		founded := uint64(1920)

		teams := []*app.Team{}

		team := app.Team{
			ID:           241,
			Name:         "Romford FC",
			ShortCode:    &short,
			CountryID:    8,
			VenueID:      9182,
			NationalTeam: false,
			Founded:      &founded,
			Logo:         nil,
		}

		teams = append(teams, &team)

		composer.On("TeamsBySeasonId", uint64(16036)).Return(teams, nil)

		handler.TeamsBySeasonId(res, req, params)

		expected := `{"status":"success","data":{"teams":[{"id":241,"name":"Romford FC","shortCode":"RFC","countryId":8,"venueId":9182,` +
			`"nationalTeam":false,"founded":1920,"logo":null}]}}`

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expected, res.Body.String())
		composer.AssertExpectations(t)
	})

	t.Run("returns a 200 response containing and empty array", func(t *testing.T) {
		t.Helper()

		composer := new(mock.TeamComposer)
		handler := rest.NewTeamHandler(composer)

		req := httptest.NewRequest(http.MethodGet, "/season/16036/teams", nil)
		res := httptest.NewRecorder()
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "16036",
			},
		}

		composer.On("TeamsBySeasonId", uint64(16036)).Return([]*app.Team{}, nil)

		handler.TeamsBySeasonId(res, req, params)

		expected := `{"status":"success","data":{"teams":[]}}`

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expected, res.Body.String())
		composer.AssertExpectations(t)
	})

	t.Run("returns a 400 response if unable to parse id request parameter", func(t *testing.T) {
		t.Helper()

		t.Helper()

		composer := new(mock.TeamComposer)
		handler := rest.NewTeamHandler(composer)

		req := httptest.NewRequest(http.MethodGet, "/season/16036/teams", nil)
		res := httptest.NewRecorder()
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "hello",
			},
		}

		handler.TeamsBySeasonId(res, req, params)

		expected := `{"status":"fail","data":[{"message":"unable to parse ID parameter as correct schema","code":1}]}`

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, expected, res.Body.String())
		composer.AssertNotCalled(t, "TeamsBySeasonId")
	})

	t.Run("returns a 500 response if internal server error returned by composer", func(t *testing.T) {
		t.Helper()

		composer := new(mock.TeamComposer)
		handler := rest.NewTeamHandler(composer)

		req := httptest.NewRequest(http.MethodGet, "/season/16036/teams", nil)
		res := httptest.NewRecorder()
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "16036",
			},
		}

		composer.On("TeamsBySeasonId", uint64(16036)).Return([]*app.Team{}, errors.ErrorInternalServerError)

		handler.TeamsBySeasonId(res, req, params)

		expected := `{"status":"error","data":[{"message":"internal server error","code":1}]}`

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, expected, res.Body.String())
		composer.AssertExpectations(t)
	})

	t.Run("returns a 502 response if bad gateway error returned by composer", func(t *testing.T) {
		t.Helper()

		composer := new(mock.TeamComposer)
		handler := rest.NewTeamHandler(composer)

		req := httptest.NewRequest(http.MethodGet, "/season/16036/teams", nil)
		res := httptest.NewRecorder()
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "16036",
			},
		}

		composer.On("TeamsBySeasonId", uint64(16036)).Return([]*app.Team{}, errors.ErrorBadGateway)

		handler.TeamsBySeasonId(res, req, params)

		expected := `{"status":"error","data":[{"message":"bad gateway","code":1}]}`

		assert.Equal(t, http.StatusBadGateway, res.Code)
		assert.Equal(t, expected, res.Body.String())
		composer.AssertExpectations(t)
	})
}
