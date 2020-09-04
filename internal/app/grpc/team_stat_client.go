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
	Stats(ctx context.Context, req *proto.TeamStatRequest) (<-chan *app.TeamStat, chan error)
}

type teamStatClient struct {
	client proto.TeamStatsServiceClient
	logger *logrus.Logger
}

func (t *teamStatClient) Stats(ctx context.Context, req *proto.TeamStatRequest) (<-chan *app.TeamStat, chan error) {
	stats := make(chan *app.TeamStat)
	errChan := make(chan error, 1)

	stream, err := t.client.GetStatForTeam(ctx, req)

	if err != nil {
		defer close(stats)
		defer close(errChan)

		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.InvalidArgument:
				errChan <- err
				break
			case codes.Internal:
				t.logError(err)
				errChan <- errors.ErrorInternalServerError
				break
			default:
				t.logError(err)
				errChan <- errors.ErrorBadGateway
			}
		}

		return stats, errChan
	}

	go t.streamStats(stream, stats, errChan)

	return stats, errChan
}

func (t *teamStatClient) streamStats(stream proto.TeamStatsService_GetStatForTeamClient, ch chan<- *app.TeamStat, errChan chan<- error) {
	for {
		stat, err := stream.Recv()

		if err != nil {
			switch err {
			case io.EOF:
				break
			default:
				t.logError(err)
				errChan <- errors.ErrorInternalServerError
			}

			close(ch)
			close(errChan)
			return
		}

		ch <- convertTeamStat(stat)
	}
}

func (t *teamStatClient) logError(err error) {
	t.logger.Errorf("Error in result client %s", err.Error())
}

func NewTeamStatClient(p proto.TeamStatsServiceClient, l *logrus.Logger) TeamStatClient {
	return &teamStatClient{client: p, logger: l}
}
