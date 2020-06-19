package proxy

import (
	"context"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc/proto"
)

type PerformanceService struct {
	client grpc.PerformanceClient
}

func (p *PerformanceService) ProxyTeamsMatchingFilterRequest(filter app.StatFilter, seasons []uint64) ([]*app.Team, error) {
	request := proto.TeamStatRequest{
		Action:  filter.Action,
		Games:   uint32(filter.Games),
		Measure: filter.Measure,
		Metric:  filter.Metric,
		Seasons: seasons,
		Stat:    filter.Stat,
		Value:   filter.Value,
		Venue:   filter.Venue,
	}

	return p.client.GetTeamsMatchingStat(context.Background(), &request)
}
