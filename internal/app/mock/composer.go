package mock

import (
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/stretchr/testify/mock"
)

type TeamComposer struct {
	mock.Mock
}

func (t *TeamComposer) TeamById(id uint64) (*app.Team, error) {
	args := t.Called(id)
	return args.Get(0).(*app.Team), args.Error(1)
}
