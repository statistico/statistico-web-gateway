package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/errors"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

type SeasonClient interface {
	ByTeamId(ctx context.Context, teamId uint64, sort string) ([]*app.Season, error)
}

type seasonClient struct {
	client proto.SeasonServiceClient
	logger *logrus.Logger
}

func (s seasonClient) ByTeamId(ctx context.Context, teamId uint64, sort string) ([]*app.Season, error) {
	seasons := []*app.Season{}

	req := proto.TeamSeasonsRequest{
		TeamId: teamId,
		Sort:   &wrappers.StringValue{Value: sort},
	}

	stream, err := s.client.GetSeasonsForTeam(ctx, &req)

	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.Internal:
				s.logError(err)
				return seasons, errors.ErrorInternalServerError
			default:
				s.logError(err)
				return seasons, errors.ErrorBadGateway
			}
		}
	}

	for {
		season, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			s.logError(err)
			return seasons, errors.ErrorInternalServerError
		}

		s := convertSeason(season)

		seasons = append(seasons, &s)
	}

	return seasons, nil
}

func (s seasonClient) logError(err error) {
	s.logger.Errorf("Error in competition client: %s", err.Error())
}

func NewSeasonClient(c proto.SeasonServiceClient, l *logrus.Logger) SeasonClient {
	return &seasonClient{
		client: c,
		logger: l,
	}
}
