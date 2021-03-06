package mock

import (
	"context"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/composer"
	"github.com/stretchr/testify/mock"
)

type CompetitionComposer struct {
	mock.Mock
}

func (c *CompetitionComposer) CompetitionsByCountryId(countryId uint64) ([]*app.Competition, error) {
	args := c.Called(countryId)
	return args.Get(0).([]*app.Competition), args.Error(1)
}

func (c *CompetitionComposer) CompetitionSeasons(competitionId uint64, sort string) ([]*app.Season, error) {
	args := c.Called(competitionId, sort)
	return args.Get(0).([]*app.Season), args.Error(1)
}

type TeamComposer struct {
	mock.Mock
}

func (t *TeamComposer) TeamById(id uint64) (*app.Team, error) {
	args := t.Called(id)
	return args.Get(0).(*app.Team), args.Error(1)
}

func (t *TeamComposer) TeamsBySeasonId(seasonId uint64) ([]*app.Team, error) {
	args := t.Called(seasonId)
	return args.Get(0).([]*app.Team), args.Error(1)
}

type ResultComposer struct {
	mock.Mock
}

func (r *ResultComposer) FetchResults(filters *composer.ResultFilters) ([]*app.Result, error) {
	args := r.Called(filters)
	return args.Get(0).([]*app.Result), args.Error(1)
}

type SeasonComposer struct {
	mock.Mock
}

func (s *SeasonComposer) ByTeamId(teamId uint64, sort string) ([]*app.Season, error) {
	args := s.Called(teamId, sort)
	return args.Get(0).([]*app.Season), args.Error(1)
}

type TeamStatComposer struct {
	mock.Mock
}

func (t *TeamStatComposer) FetchStats(ctx context.Context, filters *composer.TeamStatFilters) ([]*app.TeamStat, error) {
	args := t.Called(ctx, filters)
	return args.Get(0).([]*app.TeamStat), args.Error(1)
}
