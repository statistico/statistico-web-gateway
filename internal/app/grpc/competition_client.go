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

type CompetitionClient interface {
	CompetitionByCountryId(ctx context.Context, countryId uint64) ([]*app.Competition, error)
}

type competitionClient struct {
	client proto.CompetitionServiceClient
	logger *logrus.Logger
}

func (c competitionClient) CompetitionByCountryId(ctx context.Context, countryId uint64) ([]*app.Competition, error) {
	competitions := []*app.Competition{}

	req := proto.CompetitionRequest{CountryIds: []uint64{countryId}}

	stream, err := c.client.ListCompetitions(ctx, &req)

	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.Internal:
				c.logError(err)
				return competitions, errors.ErrorInternalServerError
			default:
				c.logError(err)
				return competitions, errors.ErrorBadGateway
			}
		}
	}

	for {
		competition, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			c.logError(err)
			return competitions, errors.ErrorInternalServerError
		}

		competitions = append(competitions, convertCompetition(competition))
	}

	return competitions, nil
}

func (c competitionClient) logError(err error) {
	c.logger.Errorf("Error in competition client: %s", err.Error())
}

func NewCompetitionClient(p proto.CompetitionServiceClient, l *logrus.Logger) CompetitionClient {
	return competitionClient{
		client: p,
		logger: l,
	}
}
