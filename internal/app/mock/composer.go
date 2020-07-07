package mock

import (
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/composer"
	"github.com/stretchr/testify/mock"
)

type TeamComposer struct {
	mock.Mock
}

func (t *TeamComposer) TeamById(id uint64) (*app.Team, error) {
	args := t.Called(id)
	return args.Get(0).(*app.Team), args.Error(1)
}

type ResultComposer struct {
	mock.Mock
}

func (r *ResultComposer) ForTeam(teamId uint64, filters *composer.Filters) ([]*app.Result, error) {
	args := r.Called(teamId, filters)
	return args.Get(0).([]*app.Result), args.Error(1)
}
