package mock

import (
	"context"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc/proto"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type CompetitionClient struct {
	mock.Mock
}

func (c *CompetitionClient) ListCompetitions(ctx context.Context, in *proto.CompetitionRequest, opts ...grpc.CallOption) (proto.CompetitionService_ListCompetitionsClient, error) {
	args := c.Called(ctx, in, opts)
	return args.Get(0).(proto.CompetitionService_ListCompetitionsClient), args.Error(1)
}

type CompetitionStream struct {
	mock.Mock
	grpc.ClientStream
}

func (c *CompetitionStream) Recv() (*proto.Competition, error) {
	args := c.Called()
	return args.Get(0).(*proto.Competition), args.Error(1)
}

type SeasonClient struct {
	mock.Mock
}

func (s *SeasonClient) GetSeasonsForCompetition(ctx context.Context, in *proto.SeasonCompetitionRequest, opts ...grpc.CallOption) (proto.SeasonService_GetSeasonsForCompetitionClient, error) {
	args := s.Called(ctx, in, opts)
	return args.Get(0).(proto.SeasonService_GetSeasonsForCompetitionClient), args.Error(1)
}

func (s *SeasonClient) GetSeasonsForTeam(ctx context.Context, in *proto.TeamSeasonsRequest, opts ...grpc.CallOption) (proto.SeasonService_GetSeasonsForTeamClient, error) {
	args := s.Called(ctx, in, opts)
	return args.Get(0).(proto.SeasonService_GetSeasonsForTeamClient), args.Error(1)
}

type SeasonStream struct {
	mock.Mock
	grpc.ClientStream
}

func (s *SeasonStream) Recv() (*proto.Season, error) {
	args := s.Called()
	return args.Get(0).(*proto.Season), args.Error(1)
}

type TeamClient struct {
	mock.Mock
}

func (t *TeamClient) GetTeamByID(ctx context.Context, in *proto.TeamRequest, opts ...grpc.CallOption) (*proto.Team, error) {
	args := t.Called(ctx, in, opts)
	return args.Get(0).(*proto.Team), args.Error(1)
}

func (t *TeamClient) GetTeamsBySeasonId(ctx context.Context, in *proto.SeasonTeamsRequest, opts ...grpc.CallOption) (proto.TeamService_GetTeamsBySeasonIdClient, error) {
	args := t.Called(ctx, in, opts)
	return args.Get(0).(proto.TeamService_GetTeamsBySeasonIdClient), args.Error(1)
}

type TeamStream struct {
	mock.Mock
	grpc.ClientStream
}

func (t *TeamStream) Recv() (*proto.Team, error) {
	args := t.Called()
	return args.Get(0).(*proto.Team), args.Error(1)
}

type ResultClient struct {
	mock.Mock
}

func (t *ResultClient) GetResultsForTeam(ctx context.Context, in *proto.TeamResultRequest, opts ...grpc.CallOption) (proto.ResultService_GetResultsForTeamClient, error) {
	args := t.Called(ctx, in, opts)
	return args.Get(0).(proto.ResultService_GetResultsForTeamClient), args.Error(1)
}

type ResultStream struct {
	mock.Mock
	grpc.ClientStream
}

func (r *ResultStream) Recv() (*proto.Result, error) {
	args := r.Called()
	return args.Get(0).(*proto.Result), args.Error(1)
}


type TeamStatClient struct {
	mock.Mock
}

func (t *TeamStatClient) GetStatForTeam(ctx context.Context, in *proto.TeamStatRequest, opts ...grpc.CallOption) (proto.TeamStatsService_GetStatForTeamClient, error) {
	args := t.Called(ctx, in, opts)
	return args.Get(0).(proto.TeamStatsService_GetStatForTeamClient), args.Error(1)
}

type TeamStatStream struct {
	mock.Mock
	grpc.ClientStream
}

func (t *TeamStatStream) Recv() (*proto.TeamStat, error) {
	args := t.Called()
	return args.Get(0).(*proto.TeamStat), args.Error(1)
}
