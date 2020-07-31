package grpc_test

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	g "github.com/statistico/statistico-web-gateway/internal/app/grpc"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc/proto"
	"github.com/statistico/statistico-web-gateway/internal/app/mock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"testing"
)

func TestTeamDataClient_TeamById(t *testing.T) {
	t.Run("calls fixture client and return team struct", func(t *testing.T) {
		t.Helper()

		m := new(mock.TeamClient)
		logger, hook := test.NewNullLogger()
		client := g.NewTeamClient(m, logger)

		request := proto.TeamRequest{TeamId: 1}

		response := proto.Team{
			Id:             1,
			Name:           "West Ham United",
			ShortCode:      &wrappers.StringValue{Value: "WHU"},
			CountryId:      8,
			VenueId:        214,
			IsNationalTeam: &wrappers.BoolValue{Value: false},
			Founded:        &wrappers.UInt64Value{Value: 1895},
			Logo:           &wrappers.StringValue{Value: "logo"},
		}

		ctx := context.Background()

		m.On("GetTeamByID", ctx, &request, []grpc.CallOption(nil)).Return(&response, nil)

		team, err := client.TeamById(ctx, &request)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		a := assert.New(t)
		a.Equal(uint64(1), team.ID)
		a.Equal("West Ham United", team.Name)
		a.Equal("WHU", *team.ShortCode)
		a.Equal(uint64(8), team.CountryID)
		a.Equal(uint64(214), team.VenueID)
		a.Equal(false, team.NationalTeam)
		a.Equal(uint64(1895), *team.Founded)
		a.Equal("logo", *team.Logo)
		a.Nil(hook.LastEntry())
		m.AssertExpectations(t)
	})

	t.Run("parses nullable fields from team returned in response", func(t *testing.T) {
		t.Helper()

		m := new(mock.TeamClient)
		logger, hook := test.NewNullLogger()
		client := g.NewTeamClient(m, logger)

		request := proto.TeamRequest{TeamId: 1}

		response := proto.Team{
			Id:        1,
			Name:      "West Ham United",
			CountryId: 8,
			VenueId:   214,
		}

		ctx := context.Background()

		m.On("GetTeamByID", ctx, &request, []grpc.CallOption(nil)).Return(&response, nil)

		team, err := client.TeamById(ctx, &request)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		a := assert.New(t)
		a.Equal(uint64(1), team.ID)
		a.Equal("West Ham United", team.Name)
		a.Nil(team.ShortCode)
		a.Equal(uint64(8), team.CountryID)
		a.Equal(uint64(214), team.VenueID)
		a.Equal(false, team.NationalTeam)
		a.Nil(team.Founded)
		a.Nil(team.Logo)
		a.Nil(hook.LastEntry())
		m.AssertExpectations(t)
	})

	t.Run("returns a not found if not found error is returned by grpc client", func(t *testing.T) {
		t.Helper()

		m := new(mock.TeamClient)
		logger, hook := test.NewNullLogger()
		client := g.NewTeamClient(m, logger)

		request := proto.TeamRequest{TeamId: 1}

		ctx := context.Background()

		e := status.Error(codes.NotFound, "not found")

		m.On("GetTeamByID", ctx, &request, []grpc.CallOption(nil)).Return(&proto.Team{}, e)

		_, err := client.TeamById(ctx, &request)

		if err == nil {
			t.Fatal("Expected errors, got nil")
		}

		assert.Equal(t, "the resource requested does not exist", err.Error())
		assert.Nil(t, hook.LastEntry())
		m.AssertExpectations(t)
	})

	t.Run("returns a bad gateway error", func(t *testing.T) {
		t.Helper()

		m := new(mock.TeamClient)
		logger, hook := test.NewNullLogger()
		client := g.NewTeamClient(m, logger)

		request := proto.TeamRequest{TeamId: 1}

		ctx := context.Background()

		e := status.Error(codes.Aborted, "aborted")

		m.On("GetTeamByID", ctx, &request, []grpc.CallOption(nil)).Return(&proto.Team{}, e)

		_, err := client.TeamById(ctx, &request)

		if err == nil {
			t.Fatal("Expected errors, got nil")
		}

		assert.Equal(t, "error response returned from external service", err.Error())
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	})

	t.Run("returns an internal error", func(t *testing.T) {
		t.Helper()

		m := new(mock.TeamClient)
		logger, hook := test.NewNullLogger()
		client := g.NewTeamClient(m, logger)

		request := proto.TeamRequest{TeamId: 1}

		ctx := context.Background()

		e := errors.New("internal server error")

		m.On("GetTeamByID", ctx, &request, []grpc.CallOption(nil)).Return(&proto.Team{}, e)

		_, err := client.TeamById(ctx, &request)

		if err == nil {
			t.Fatal("Expected errors, got nil")
		}

		assert.Equal(t, "internal server error", err.Error())
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	})
}

func TestTeamClient_TeamsBySeasonId(t *testing.T) {
	t.Run("calls team client and returns a slice of team struct", func(t *testing.T) {
		t.Helper()

		m := new(mock.TeamClient)
		logger, hook := test.NewNullLogger()
		client := g.NewTeamClient(m, logger)

		stream := new(mock.TeamStream)

		team := proto.Team{
			Id:        1,
			Name:      "West Ham United",
			CountryId: 8,
			VenueId:   214,
		}

		ctx := context.Background()

		request := proto.SeasonTeamsRequest{SeasonId: 16036}

		m.On("GetTeamsBySeasonId", ctx, &request, []grpc.CallOption(nil)).Return(stream, nil)
		stream.On("Recv").Twice().Return(&team, nil)
		stream.On("Recv").Once().Return(&proto.Team{}, io.EOF)

		teams, err := client.TeamsBySeasonId(ctx, 16036)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Equal(t, 2, len(teams))
		assert.Nil(t, hook.LastEntry())
		m.AssertExpectations(t)
		stream.AssertExpectations(t)
	})

	t.Run("logs error and returns internal server error if internal server error is returned by client", func(t *testing.T) {
		t.Helper()

		m := new(mock.TeamClient)
		logger, hook := test.NewNullLogger()
		client := g.NewTeamClient(m, logger)

		stream := new(mock.TeamStream)

		ctx := context.Background()

		request := proto.SeasonTeamsRequest{SeasonId: 16036}

		e := status.Error(codes.Internal, "internal error")

		m.On("GetTeamsBySeasonId", ctx, &request, []grpc.CallOption(nil)).Return(stream, e)

		_, err := client.TeamsBySeasonId(ctx, 16036)

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "internal server error", err.Error())
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		assert.Equal(t, "Error in team client: rpc error: code = Internal desc = internal error", hook.LastEntry().Message)
		m.AssertExpectations(t)
		stream.AssertNotCalled(t, "Recv")
	})

	t.Run("logs error and returns bad gateway error for non internal server error returned by client", func(t *testing.T) {
		t.Helper()

		m := new(mock.TeamClient)
		logger, hook := test.NewNullLogger()
		client := g.NewTeamClient(m, logger)

		stream := new(mock.TeamStream)

		ctx := context.Background()

		request := proto.SeasonTeamsRequest{SeasonId: 16036}

		e := status.Error(codes.Unavailable, "service unavailable")

		m.On("GetTeamsBySeasonId", ctx, &request, []grpc.CallOption(nil)).Return(stream, e)

		_, err := client.TeamsBySeasonId(ctx, 16036)

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "error response returned from external service", err.Error())
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		assert.Equal(t, "Error in team client: rpc error: code = Unavailable desc = service unavailable", hook.LastEntry().Message)
		m.AssertExpectations(t)
		stream.AssertNotCalled(t, "Recv")
	})

	t.Run("logs error and returns internal server error if error reading from stream", func(t *testing.T) {
		t.Helper()

		m := new(mock.TeamClient)
		logger, hook := test.NewNullLogger()
		client := g.NewTeamClient(m, logger)

		stream := new(mock.TeamStream)

		team := proto.Team{
			Id:        1,
			Name:      "West Ham United",
			CountryId: 8,
			VenueId:   214,
		}

		ctx := context.Background()

		request := proto.SeasonTeamsRequest{SeasonId: 16036}

		e := errors.New("oh damn")

		m.On("GetTeamsBySeasonId", ctx, &request, []grpc.CallOption(nil)).Return(stream, nil)
		stream.On("Recv").Twice().Return(&team, nil)
		stream.On("Recv").Once().Return(&proto.Team{}, e)

		_, err := client.TeamsBySeasonId(ctx, 16036)

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "internal server error", err.Error())
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		assert.Equal(t, "Error in team client: oh damn", hook.LastEntry().Message)
		m.AssertExpectations(t)
		stream.AssertExpectations(t)
	})
}