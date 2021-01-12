package composer

import (
	"context"
	"github.com/statistico/statistico-data-go-grpc-client"
	"github.com/statistico/statistico-proto/go"
)

type TeamComposer interface {
	ByID(id uint64) (*statisticoproto.Team, error)
	BySeasonID(seasonId uint64) ([]*statisticoproto.Team, error)
}

type teamComposer struct {
	client statisticodata.TeamClient
}

func (t *teamComposer) ByID(id uint64) (*statisticoproto.Team, error) {
	return t.client.ByID(context.Background(), id)
}

func (t teamComposer) BySeasonID(seasonId uint64) ([]*statisticoproto.Team, error) {
	return t.client.BySeasonID(context.Background(), seasonId)
}

func NewTeamComposer(c statisticodata.TeamClient) TeamComposer {
	return &teamComposer{client: c}
}
