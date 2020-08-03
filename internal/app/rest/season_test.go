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

func TestSeasonHandler_ByCompetitionId(t *testing.T) {
	t.Run("returns a 200 response containing seasons data", func(t *testing.T) {
		t.Helper()

		c := new(mock.CompetitionComposer)
		s := new(mock.SeasonComposer)
		handler := rest.NewSeasonHandler(c, s)

		req := httptest.NewRequest(http.MethodGet, "/competition/8/seasons", nil)
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "8",
			},
		}

		res := httptest.NewRecorder()

		c.On("CompetitionSeasons", uint64(8), "").Return(seasons(), nil)

		handler.ByCompetitionId(res, req, params)

		expected := `{"status":"success","data":{"seasons":[{"id":12968,"name":"2018/2019","isCurrent":false},` +
			`{"id":16036,"name":"2019/2020","isCurrent":true}]}}`

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})

	t.Run("parses sort query parameter and returns a 200 response containing seasons data", func(t *testing.T) {
		t.Helper()

		c := new(mock.CompetitionComposer)
		s := new(mock.SeasonComposer)
		handler := rest.NewSeasonHandler(c, s)

		req := httptest.NewRequest(http.MethodGet, "/competition/8/seasons", nil)

		q := req.URL.Query()
		q.Add("sort", "name_asc")
		req.URL.RawQuery = q.Encode()

		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "8",
			},
		}

		res := httptest.NewRecorder()

		c.On("CompetitionSeasons", uint64(8), "name_asc").Return(seasons(), nil)

		handler.ByCompetitionId(res, req, params)

		expected := `{"status":"success","data":{"seasons":[{"id":12968,"name":"2018/2019","isCurrent":false},` +
			`{"id":16036,"name":"2019/2020","isCurrent":true}]}}`

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})

	t.Run("returns a 200 response containing an empty array", func(t *testing.T) {
		t.Helper()

		c := new(mock.CompetitionComposer)
		s := new(mock.SeasonComposer)
		handler := rest.NewSeasonHandler(c, s)

		req := httptest.NewRequest(http.MethodGet, "/competition/8/seasons", nil)
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "8",
			},
		}

		res := httptest.NewRecorder()

		c.On("CompetitionSeasons", uint64(8), "").Return([]*app.Season{}, nil)

		handler.ByCompetitionId(res, req, params)

		expected := `{"status":"success","data":{"seasons":[]}}`

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})

	t.Run("returns a 400 response if unable to parse id request parameter", func(t *testing.T) {
		t.Helper()

		c := new(mock.CompetitionComposer)
		s := new(mock.SeasonComposer)
		handler := rest.NewSeasonHandler(c, s)

		req := httptest.NewRequest(http.MethodGet, "/competition/8/seasons", nil)
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "hello",
			},
		}

		res := httptest.NewRecorder()

		handler.ByCompetitionId(res, req, params)

		expected := `{"status":"fail","data":[{"message":"unable to parse ID parameter as correct schema","code":1}]}`

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertNotCalled(t, "CompetitionSeasons")
	})

	t.Run("returns a 500 response if internal server error returned by composer", func(t *testing.T) {
		t.Helper()

		c := new(mock.CompetitionComposer)
		s := new(mock.SeasonComposer)
		handler := rest.NewSeasonHandler(c, s)

		req := httptest.NewRequest(http.MethodGet, "/competition/8/seasons", nil)
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "8",
			},
		}

		res := httptest.NewRecorder()

		c.On("CompetitionSeasons", uint64(8), "").Return([]*app.Season{}, e.ErrorInternalServerError)

		handler.ByCompetitionId(res, req, params)

		expected := `{"status":"error","data":[{"message":"internal server error","code":1}]}`

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})

	t.Run("returns a 502 response if bad gateway error returned by composer", func(t *testing.T) {
		t.Helper()

		c := new(mock.CompetitionComposer)
		s := new(mock.SeasonComposer)
		handler := rest.NewSeasonHandler(c, s)

		req := httptest.NewRequest(http.MethodGet, "/competition/8/seasons", nil)
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "8",
			},
		}

		res := httptest.NewRecorder()

		c.On("CompetitionSeasons", uint64(8), "").Return([]*app.Season{}, e.ErrorBadGateway)

		handler.ByCompetitionId(res, req, params)

		expected := `{"status":"error","data":[{"message":"bad gateway","code":1}]}`

		assert.Equal(t, http.StatusBadGateway, res.Code)
		assert.Equal(t, expected, res.Body.String())
		c.AssertExpectations(t)
	})
}

func TestSeasonHandler_ByTeamId(t *testing.T) {
	t.Run("returns a 200 response containing seasons data", func(t *testing.T) {
		t.Helper()

		c := new(mock.CompetitionComposer)
		s := new(mock.SeasonComposer)
		handler := rest.NewSeasonHandler(c, s)

		req := httptest.NewRequest(http.MethodGet, "/team/8/seasons", nil)
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "8",
			},
		}

		res := httptest.NewRecorder()

		s.On("ByTeamId", uint64(8), "").Return(seasons(), nil)

		handler.ByTeamId(res, req, params)

		expected := `{"status":"success","data":{"seasons":[{"id":12968,"name":"2018/2019","isCurrent":false},` +
			`{"id":16036,"name":"2019/2020","isCurrent":true}]}}`

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expected, res.Body.String())
		s.AssertExpectations(t)
	})

	t.Run("parses sort query parameter and returns a 200 response containing seasons data", func(t *testing.T) {
		t.Helper()

		c := new(mock.CompetitionComposer)
		s := new(mock.SeasonComposer)
		handler := rest.NewSeasonHandler(c, s)

		req := httptest.NewRequest(http.MethodGet, "/team/8/seasons", nil)

		q := req.URL.Query()
		q.Add("sort", "name_asc")
		req.URL.RawQuery = q.Encode()

		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "8",
			},
		}

		res := httptest.NewRecorder()

		s.On("ByTeamId", uint64(8), "name_asc").Return(seasons(), nil)

		handler.ByTeamId(res, req, params)

		expected := `{"status":"success","data":{"seasons":[{"id":12968,"name":"2018/2019","isCurrent":false},` +
			`{"id":16036,"name":"2019/2020","isCurrent":true}]}}`

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expected, res.Body.String())
		s.AssertExpectations(t)
	})

	t.Run("returns a 200 response containing an empty array", func(t *testing.T) {
		t.Helper()

		c := new(mock.CompetitionComposer)
		s := new(mock.SeasonComposer)
		handler := rest.NewSeasonHandler(c, s)

		req := httptest.NewRequest(http.MethodGet, "/team/8/seasons", nil)
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "8",
			},
		}

		res := httptest.NewRecorder()

		s.On("ByTeamId", uint64(8), "").Return([]*app.Season{}, nil)

		handler.ByTeamId(res, req, params)

		expected := `{"status":"success","data":{"seasons":[]}}`

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expected, res.Body.String())
		s.AssertExpectations(t)
	})

	t.Run("returns a 400 response if unable to parse id request parameter", func(t *testing.T) {
		t.Helper()

		c := new(mock.CompetitionComposer)
		s := new(mock.SeasonComposer)
		handler := rest.NewSeasonHandler(c, s)

		req := httptest.NewRequest(http.MethodGet, "/team/8/seasons", nil)
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "hello",
			},
		}

		res := httptest.NewRecorder()

		handler.ByTeamId(res, req, params)

		expected := `{"status":"fail","data":[{"message":"unable to parse ID parameter as correct schema","code":1}]}`

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, expected, res.Body.String())
		s.AssertNotCalled(t, "ByTeamId")
	})

	t.Run("returns a 500 response if internal server error returned by composer", func(t *testing.T) {
		t.Helper()

		c := new(mock.CompetitionComposer)
		s := new(mock.SeasonComposer)
		handler := rest.NewSeasonHandler(c, s)

		req := httptest.NewRequest(http.MethodGet, "/team/8/seasons", nil)
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "8",
			},
		}

		res := httptest.NewRecorder()

		s.On("ByTeamId", uint64(8), "").Return([]*app.Season{}, e.ErrorInternalServerError)

		handler.ByTeamId(res, req, params)

		expected := `{"status":"error","data":[{"message":"internal server error","code":1}]}`

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, expected, res.Body.String())
		s.AssertExpectations(t)
	})

	t.Run("returns a 502 response if bad gateway error returned by composer", func(t *testing.T) {
		t.Helper()

		c := new(mock.CompetitionComposer)
		s := new(mock.SeasonComposer)
		handler := rest.NewSeasonHandler(c, s)

		req := httptest.NewRequest(http.MethodGet, "/team/8/seasons", nil)
		params := httprouter.Params{
			httprouter.Param{
				Key:   "id",
				Value: "8",
			},
		}
	
		res := httptest.NewRecorder()

		s.On("ByTeamId", uint64(8), "").Return([]*app.Season{}, e.ErrorBadGateway)

		handler.ByTeamId(res, req, params)

		expected := `{"status":"error","data":[{"message":"bad gateway","code":1}]}`

		assert.Equal(t, http.StatusBadGateway, res.Code)
		assert.Equal(t, expected, res.Body.String())
		s.AssertExpectations(t)
	})
}

func seasons() []*app.Season {
	one := app.Season{
		ID:        12968,
		Name:      "2018/2019",
		IsCurrent: false,
	}

	two := app.Season{
		ID:        16036,
		Name:      "2019/2020",
		IsCurrent: true,
	}

	return []*app.Season{&one, &two}
}