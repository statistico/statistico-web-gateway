package rest_test

import (
	"github.com/julienschmidt/httprouter"
	"github.com/statistico/statistico-web-gateway/internal/app"
	e "github.com/statistico/statistico-web-gateway/internal/app/errors"
	"github.com/statistico/statistico-web-gateway/internal/app/mock"
	"github.com/statistico/statistico-web-gateway/internal/app/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCompetitionHandler_ByCountryId(t *testing.T) {
	t.Run("returns a 200 response containing competitions data", func(t *testing.T) {
		t.Helper()

		c := new(mock.CompetitionComposer)
		handler := rest.NewCompetitionHandler(c)

		req := httptest.NewRequest(http.MethodGet, "/country/462/competitions", nil)
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "462",
			},
		}

		res := httptest.NewRecorder()

		c.On("CompetitionsByCountryId", uint64(462)).Return(competitions(), nil)

		handler.ByCountryId(res, req, params)

		expected := `{"status":"success","data":{"competitions":[{"id":8,"name":"Premier League","isCup":false,"countryId":462},` +
			`{"id":10,"name":"Championship","isCup":false,"countryId":462}]}}`

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})

	t.Run("returns a 200 response containing an empty array", func(t *testing.T) {
		t.Helper()

		c := new(mock.CompetitionComposer)
		handler := rest.NewCompetitionHandler(c)

		req := httptest.NewRequest(http.MethodGet, "/country/462/competitions", nil)
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "462",
			},
		}

		res := httptest.NewRecorder()

		c.On("CompetitionsByCountryId", uint64(462)).Return([]*app.Competition{}, nil)

		handler.ByCountryId(res, req, params)

		expected := `{"status":"success","data":{"competitions":[]}}`

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})

	t.Run("returns a 400 response if unable to parse id request parameter", func(t *testing.T) {
		t.Helper()

		c := new(mock.CompetitionComposer)
		handler := rest.NewCompetitionHandler(c)

		req := httptest.NewRequest(http.MethodGet, "/country/462/competitions", nil)
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "hello",
			},
		}

		res := httptest.NewRecorder()

		handler.ByCountryId(res, req, params)

		expected := `{"status":"fail","data":[{"message":"unable to parse ID parameter as correct schema","code":1}]}`

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertNotCalled(t, "CompetitionsByCountryId")
	})

	t.Run("returns a 500 response if internal server error returned by composer", func(t *testing.T) {
		t.Helper()

		c := new(mock.CompetitionComposer)
		handler := rest.NewCompetitionHandler(c)

		req := httptest.NewRequest(http.MethodGet, "/country/462/competitions", nil)
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "462",
			},
		}

		res := httptest.NewRecorder()

		c.On("CompetitionsByCountryId", uint64(462)).Return([]*app.Competition{}, e.ErrorInternalServerError)

		expected := `{"status":"error","data":[{"message":"internal server error","code":1}]}`

		handler.ByCountryId(res, req, params)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})

	t.Run("returns a 500 response if internal server error returned by composer", func(t *testing.T) {
		t.Helper()

		c := new(mock.CompetitionComposer)
		handler := rest.NewCompetitionHandler(c)

		req := httptest.NewRequest(http.MethodGet, "/country/462/competitions", nil)
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "462",
			},
		}

		res := httptest.NewRecorder()

		c.On("CompetitionsByCountryId", uint64(462)).Return([]*app.Competition{}, e.ErrorBadGateway)

		expected := `{"status":"error","data":[{"message":"bad gateway","code":1}]}`

		handler.ByCountryId(res, req, params)

		assert.Equal(t, http.StatusBadGateway, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})
}

func competitions() []*app.Competition {
	one := app.Competition{
		ID:        8,
		Name:      "Premier League",
		IsCup:     false,
		CountryID: 462,
	}

	two := app.Competition{
		ID:        10,
		Name:      "Championship",
		IsCup:     false,
		CountryID: 462,
	}

	return []*app.Competition{&one, &two}
}