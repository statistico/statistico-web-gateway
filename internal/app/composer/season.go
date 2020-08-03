package composer

import (
	"context"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc"
)

type SeasonComposer interface {
	ByTeamId(teamId uint64, sort string) ([]*app.Season, error)
}

type seasonComposer struct {
	client grpc.SeasonClient
}

func (s *seasonComposer) ByTeamId(teamId uint64, sort string) ([]*app.Season, error) {
	return s.client.ByTeamId(context.Background(), teamId, sort)
}

func NewSeasonComposer(c grpc.SeasonClient) SeasonComposer {
	return &seasonComposer{client: c}
}
