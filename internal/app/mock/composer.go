package mock

import (
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/composer"
	"github.com/stretchr/testify/mock"
)

type FixtureSearch struct {
	mock.Mock
}

func (f *FixtureSearch) Search(q *composer.FixtureSearchQuery) []*app.Fixture {
	args := f.Called(q)
	return args.Get(0).([]*app.Fixture)
}
