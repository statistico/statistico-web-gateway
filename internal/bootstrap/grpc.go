package bootstrap

import "github.com/statistico/statistico-web-gateway/internal/app/grpc"

func (c Container) GRPCTeamClient() grpc.TeamClient {
	return grpc.NewTeamClient(c.StatisticoDataServiceClient, c.Logger)
}
