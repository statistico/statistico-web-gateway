package grpc

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/errors"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TeamClient interface {
	TeamById(ctx context.Context, req *proto.TeamRequest) (*app.Team, error)
}

type teamClient struct {
	client proto.TeamServiceClient
	logger *logrus.Logger
}

func (t teamClient) TeamById(ctx context.Context, req *proto.TeamRequest) (*app.Team, error) {
	response, err := t.client.GetTeamByID(ctx, req)

	if err != nil {
		t.logger.Warnf("Error in team client %s", err.Error())

		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				return nil, errors.ErrorNotFound
			default:
				return nil, errors.ErrorBadGateway
			}
		}

		return nil, errors.ErrorInternalServerError
	}

	team := app.Team{
		ID:        response.GetId(),
		Name:      response.GetName(),
		CountryID: response.GetCountryId(),
		VenueID:   response.GetVenueId(),
	}

	if response.GetShortCode() != nil {
		team.ShortCode = &response.GetShortCode().Value
	}

	if response.GetIsNationalTeam() != nil {
		team.NationalTeam = response.GetIsNationalTeam().Value
	}

	if response.GetFounded() != nil {
		team.Founded = &response.GetFounded().Value
	}

	if response.GetLogo() != nil {
		team.Logo = &response.GetLogo().Value
	}

	return &team, nil
}

func NewTeamClient(p proto.TeamServiceClient, log *logrus.Logger) TeamClient {
	return teamClient{client: p, logger: log}
}
