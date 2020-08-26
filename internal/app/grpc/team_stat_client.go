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

type TeamStatClient interface {
	Stats(ctx context.Context, req *proto.TeamStatRequest) ([]*app.TeamStat, error)
}

type teamStatClient struct {
	client proto.TeamStatsServiceClient
	logger *logrus.Logger
}

func (t *teamStatClient) Stats(ctx context.Context, req *proto.TeamStatRequest) ([]*app.TeamStat, error) {
	stats := []*app.TeamStat{}

	stream, err := t.client.GetStatForTeam(ctx, req)

	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.InvalidArgument:
				return stats, err
			case codes.Internal:
				t.logError(err)
				return stats, errors.ErrorInternalServerError
			default:
				t.logError(err)
				return stats, errors.ErrorBadGateway
			}
		}
	}

	for {
		stat, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			t.logError(err)
			return stats, errors.ErrorInternalServerError
		}

		st := convertTeamStat(stat)

		stats = append(stats, &st)
	}

	return stats, nil
}

func (t *teamStatClient) logError(err error) {
	t.logger.Errorf("Error in result client %s", err.Error())
}

func NewTeamStatClient(p proto.TeamStatsServiceClient, l *logrus.Logger) TeamStatClient {
	return &teamStatClient{client: p, logger: l}
}
