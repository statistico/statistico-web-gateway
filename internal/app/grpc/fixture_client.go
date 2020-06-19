package grpc

import (
	"context"
	"errors"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc/proto"
	"io"
	"time"
)

type FixtureClient interface {
	GetSeasonFixtures(ctx context.Context, request *proto.SeasonFixtureRequest) ([]*app.Fixture, error)
}

type fixtureClient struct {
	client proto.FixtureServiceClient
}

func (f fixtureClient) GetSeasonFixtures(ctx context.Context, request *proto.SeasonFixtureRequest) ([]*app.Fixture, error) {
	var fixtures []*app.Fixture

	stream, err := f.client.ListSeasonFixtures(context.Background(), request)

	if err != nil {
		return nil, err
	}

	for {
		fixture, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		f, err := convertFixture(fixture)

		if err != nil {
			return nil, err
		}

		fixtures = append(fixtures, f)
	}

	return fixtures, nil
}

func convertFixture(f *proto.Fixture) (*app.Fixture, error) {
	start, err := time.Parse(time.RFC3339, f.Round.StartDate)

	if err != nil {
		return nil, errors.New("date provided in request is not a valid RFC3339 formatted date")
	}

	end, err := time.Parse(time.RFC3339, f.Round.EndDate)

	if err != nil {
		return nil, errors.New("date provided in request is not a valid RFC3339 formatted date")
	}

	date, err := time.Parse(time.RFC3339, f.DateTime.Rfc)

	if err != nil {
		return nil, errors.New("date provided in request is not a valid RFC3339 formatted date")
	}

	fix := app.Fixture{
		ID:          uint64(f.Id),
		Competition: app.Competition{
			ID: uint64(f.Competition.Id),
			Name: f.Competition.Name,
			IsCup: *f.Competition.IsCup,
		},
		Season:      app.Season{
			ID: uint64(f.Season.Id),
			Name: f.Season.Name,
			IsCurrent: *f.Season.IsCurrent,
		},
		Round:       app.Round{
			ID: uint64(f.Round.Id),
			Name: f.Round.Name,
			SeasonID: uint64(f.Round.SeasonId),
			StartDate: app.JsonDate(start),
			EndDate: app.JsonDate(end),
		},
		HomeTeam:    app.Team{
			ID: uint64(f.HomeTeam.Id),
			Name: f.HomeTeam.Name,
		},
		AwayTeam:    app.Team{
			ID: uint64(f.AwayTeam.Id),
			Name: f.AwayTeam.Name,
		},
		Venue:       app.Venue{
			ID: uint64(*f.Venue.Id),
			Name: *f.Venue.Name,
		},
		Date:        app.JsonDate(date),
	}

	return &fix, nil
}

