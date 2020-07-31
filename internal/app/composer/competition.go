package composer

import (
	"context"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc"
)

type CompetitionComposer interface {
	CompetitionsByCountryId(countryId uint64) ([]*app.Competition, error)
	CompetitionSeasons(competitionId uint64, sort string) ([]*app.Season, error)
}

type competitionComposer struct {
	client grpc.CompetitionClient
}

func (c *competitionComposer) CompetitionsByCountryId(countryId uint64) ([]*app.Competition, error) {
	return c.client.CompetitionsByCountryId(context.Background(), countryId)
}

func (c *competitionComposer) CompetitionSeasons(competitionId uint64, sort string) ([]*app.Season, error) {
	return c.client.CompetitionSeasons(context.Background(), competitionId, sort)
}

func NewCompetitionComposer(c grpc.CompetitionClient) CompetitionComposer {
	return &competitionComposer{client: c}
}
