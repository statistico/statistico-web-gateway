package grpc

import (
	"context"
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
}

func (t teamClient) TeamById(ctx context.Context, req *proto.TeamRequest) (*app.Team, error) {
	response, err := t.client.GetTeamByID(ctx, req)

	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				return nil, errors.ErrNotFound
			default:
				return nil, errors.ErrorBadGateway
			}
		} else {
			return nil, errors.ErrInternalServerError
		}
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

func NewTeamClient(p proto.TeamServiceClient) TeamClient {
	return teamClient{client: p}
}
