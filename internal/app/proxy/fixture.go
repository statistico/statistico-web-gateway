package proxy

import (
	"context"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc/proto"
	"time"
)

type FixtureService struct {
	client grpc.FixtureClient
}

func (f *FixtureService) FixtureForSeasonBetween(id uint64, from time.Time, to time.Time) ([]*app.Fixture, error) {
	request := proto.SeasonFixtureRequest{
		SeasonId: id,
		DateFrom: from.Format(time.RFC3339),
		DateTo:   to.Format(time.RFC3339),
	}

	return f.client.GetSeasonFixtures(context.Background(), &request)
}
