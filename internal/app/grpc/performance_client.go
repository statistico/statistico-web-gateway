package grpc

import (
	"context"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc/proto"
)

type PerformanceClient interface {
	GetTeamsMatchingStat(ctx context.Context, request *proto.TeamStatRequest) ([]*app.Team, error)
}

type performanceClient struct {
	client proto.PerformanceServiceClient
}

func (p performanceClient) GetTeamsMatchingStat(ctx context.Context, request *proto.TeamStatRequest) ([]*app.Team, error) {
	response, err := p.client.GetTeamsMatchingStat(ctx, request)

	if err != nil {
		return nil, err
	}

	return convertTeams(response.GetTeams()), nil
}

func convertTeams(t []*proto.Team) []*app.Team {
	var teams []*app.Team

	for _, team := range t {
		x := app.Team{
			ID:   uint64(team.Id),
			Name: team.Name,
		}

		teams = append(teams, &x)
	}

	return teams
}
