package composer

import (
	"context"
	"github.com/statistico/statistico-data-go-grpc-client"
	"github.com/statistico/statistico-proto/go"
)

type SeasonComposer interface {
	ByCompetitionID(teamId uint64, sort string) ([]*statisticoproto.Season, error)
	ByTeamID(teamId uint64, sort string) ([]*statisticoproto.Season, error)
}

type seasonComposer struct {
	client statisticodata.SeasonClient
}

func (s *seasonComposer) ByCompetitionID(teamId uint64, sort string) ([]*statisticoproto.Season, error) {
	return s.client.ByCompetitionID(context.Background(), teamId, sort)
}

func (s *seasonComposer) ByTeamID(teamId uint64, sort string) ([]*statisticoproto.Season, error) {
	return s.client.ByTeamID(context.Background(), teamId, sort)
}

func NewSeasonComposer(c statisticodata.SeasonClient) SeasonComposer {
	return &seasonComposer{client: c}
}
