package grpc

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/errors"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

type TeamClient interface {
	TeamById(ctx context.Context, req *proto.TeamRequest) (*app.Team, error)
	TeamsBySeasonId(ctx context.Context, seasonId uint64) ([]*app.Team, error)
}

type teamClient struct {
	client proto.TeamServiceClient
	logger *logrus.Logger
}

func (t teamClient) TeamById(ctx context.Context, req *proto.TeamRequest) (*app.Team, error) {
	response, err := t.client.GetTeamByID(ctx, req)

	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				return nil, errors.ErrorNotFound
			default:
				t.logError(err)
				return nil, errors.ErrorBadGateway
			}
		}

		t.logError(err)

		return nil, errors.ErrorInternalServerError
	}

	team := convertTeam(response)

	return &team, nil
}

func (t teamClient) TeamsBySeasonId(ctx context.Context, seasonId uint64) ([]*app.Team, error) {
	teams := []*app.Team{}

	req := proto.SeasonTeamsRequest{SeasonId: seasonId}

	stream, err := t.client.GetTeamsBySeasonId(ctx, &req)

	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.Internal:
				t.logError(err)
				return teams, errors.ErrorInternalServerError
			default:
				t.logError(err)
				return teams, errors.ErrorBadGateway
			}
		}
	}

	for {
		team, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			t.logError(err)
			return teams, errors.ErrorInternalServerError
		}

		te := convertTeam(team)

		teams = append(teams, &te)
	}

	return teams, nil
}

func (t teamClient) logError(err error) {
	t.logger.Errorf("Error in team client: %s", err.Error())
}

func NewTeamClient(p proto.TeamServiceClient, log *logrus.Logger) TeamClient {
	return teamClient{client: p, logger: log}
}
