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

type CompetitionClient interface {
	CompetitionsByCountryId(ctx context.Context, countryId uint64) ([]*app.Competition, error)
	CompetitionSeasons(ctx context.Context, competitionId uint64, sort string) ([]*app.Season, error)
}

type competitionClient struct {
	competitionClient proto.CompetitionServiceClient
	seasonClient      proto.SeasonServiceClient
	logger            *logrus.Logger
}

func (c *competitionClient) CompetitionsByCountryId(ctx context.Context, countryId uint64) ([]*app.Competition, error) {
	competitions := []*app.Competition{}

	req := proto.CompetitionRequest{CountryIds: []uint64{countryId}}

	stream, err := c.competitionClient.ListCompetitions(ctx, &req)

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

func (c *competitionClient) CompetitionSeasons(ctx context.Context, competitionId uint64, sort string) ([]*app.Season, error) {
	seasons := []*app.Season{}

	req := proto.SeasonCompetitionRequest{CompetitionId: competitionId, Sort: &wrappers.StringValue{Value: sort}}

	stream, err := c.seasonClient.GetSeasonsForCompetition(ctx, &req)

	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.Internal:
				c.logError(err)
				return seasons, errors.ErrorInternalServerError
			default:
				c.logError(err)
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
			c.logError(err)
			return seasons, errors.ErrorInternalServerError
		}

		s := convertSeason(season)

		seasons = append(seasons, &s)
	}

	return seasons, nil
}

func (c competitionClient) logError(err error) {
	c.logger.Errorf("Error in competition client: %s", err.Error())
}

func NewCompetitionClient(c proto.CompetitionServiceClient, s proto.SeasonServiceClient, l *logrus.Logger) CompetitionClient {
	return &competitionClient{
		competitionClient: c,
		seasonClient:      s,
		logger:            l,
	}
}
