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
