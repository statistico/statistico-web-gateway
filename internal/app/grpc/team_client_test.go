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
