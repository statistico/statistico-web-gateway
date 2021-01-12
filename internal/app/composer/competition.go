package composer

import (
	"context"
	"github.com/statistico/statistico-data-go-grpc-client"
	"github.com/statistico/statistico-proto/go"
)

type CompetitionComposer interface {
	CompetitionsByCountryID(countryId uint64) ([]*statisticoproto.Competition, error)
}

type competitionComposer struct {
	client statisticodata.CompetitionClient
}

func (c *competitionComposer) CompetitionsByCountryID(countryId uint64) ([]*statisticoproto.Competition, error) {
	return c.client.ByCountryID(context.Background(), countryId)
}

func NewCompetitionComposer(c statisticodata.CompetitionClient) CompetitionComposer {
	return &competitionComposer{client: c}
}
