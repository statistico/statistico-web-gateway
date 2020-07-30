package composer

import (
	"context"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc"
)

type CompetitionComposer interface {
	CompetitionsByCountryId(countryId uint64) ([]*app.Competition, error)
}

type competitionComposer struct {
	client grpc.CompetitionClient
}

func (c *competitionComposer) CompetitionsByCountryId(countryId uint64) ([]*app.Competition, error) {
	return c.client.CompetitionsByCountryId(context.Background(), countryId)
}

func NewCompetitionComposer(c grpc.CompetitionClient) CompetitionComposer {
	return &competitionComposer{client: c}
}
