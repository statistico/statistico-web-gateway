package bootstrap

import (
	"github.com/statistico/statistico-web-gateway/internal/app/grpc"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc/proto"
)

func (c Container) GRPCCompetitionClient() grpc.CompetitionClient {
	competitionClient := proto.NewCompetitionServiceClient(c.StatisticoDataServiceConnection)
	seasonClient := proto.NewSeasonServiceClient(c.StatisticoDataServiceConnection)
	return grpc.NewCompetitionClient(competitionClient, seasonClient, c.Logger)
}

func (c Container) GRPCResultClient() grpc.ResultClient {
	client := proto.NewResultServiceClient(c.StatisticoDataServiceConnection)
	return grpc.NewResultClient(client, c.Logger)
}

func (c Container) GRPCSeasonClient() grpc.SeasonClient {
	client := proto.NewSeasonServiceClient(c.StatisticoDataServiceConnection)
	return grpc.NewSeasonClient(client, c.Logger)
}

func (c Container) GRPCTeamClient() grpc.TeamClient {
	client := proto.NewTeamServiceClient(c.StatisticoDataServiceConnection)
	return grpc.NewTeamClient(client, c.Logger)
}
