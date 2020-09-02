package grpc_test

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/statistico/statistico-web-gateway/internal/app"
	er "github.com/statistico/statistico-web-gateway/internal/app/errors"
	g "github.com/statistico/statistico-web-gateway/internal/app/grpc"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc/proto"
	"github.com/statistico/statistico-web-gateway/internal/app/mock"
	"github.com/stretchr/testify/assert"
	mock2 "github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"testing"
	"time"
)

func TestResultClient_ByTeam(t *testing.T) {
	t.Run("calls result client and returns a slice of result struct", func(t *testing.T) {
		t.Helper()

		m := new(mock.ResultClient)
		logger, hook := test.NewNullLogger()
		client := g.NewResultClient(m, logger)

		stream := new(mock.ResultStream)

		request := proto.TeamResultRequest{
			TeamId: 1,
			Limit:  &wrappers.UInt64Value{Value: 8},
		}

		ctx := context.Background()

		m.On("GetResultsForTeam", ctx, &request, []grpc.CallOption(nil)).Return(stream, nil)
		stream.On("Recv").Twice().Return(newProtoResult(), nil)
		stream.On("Recv").Once().Return(&proto.Result{}, io.EOF)

		results, err := client.ByTeam(ctx, &request)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Equal(t, 2, len(results))
		assert.Nil(t, hook.LastEntry())
		m.AssertExpectations(t)
	})

	t.Run("returns error if invalid argument error returned by result client", func(t *testing.T) {
		t.Helper()

		m := new(mock.ResultClient)
		logger, hook := test.NewNullLogger()
		client := g.NewResultClient(m, logger)

		stream := new(mock.ResultStream)

		request := proto.TeamResultRequest{
			TeamId: 1,
			Limit:  &wrappers.UInt64Value{Value: 8},
		}

		ctx := context.Background()

		e := status.Error(codes.InvalidArgument, "incorrect format")

		m.On("GetResultsForTeam", ctx, &request, []grpc.CallOption(nil)).Return(stream, e)

		_, err := client.ByTeam(ctx, &request)

		if err == nil {
			t.Fatal("Expected errors, got nil")
		}

		assert.Equal(t, "rpc error: code = InvalidArgument desc = incorrect format", err.Error())
		assert.Nil(t, hook.LastEntry())
		m.AssertExpectations(t)
	})

	t.Run("logs error and returns internal server error", func(t *testing.T) {
		t.Helper()

		m := new(mock.ResultClient)
		logger, hook := test.NewNullLogger()
		client := g.NewResultClient(m, logger)

		stream := new(mock.ResultStream)

		request := proto.TeamResultRequest{
			TeamId: 1,
			Limit:  &wrappers.UInt64Value{Value: 8},
		}

		ctx := context.Background()

		e := status.Error(codes.Internal, "internal error")

		m.On("GetResultsForTeam", ctx, &request, []grpc.CallOption(nil)).Return(stream, e)

		_, err := client.ByTeam(ctx, &request)

		if err == nil {
			t.Fatal("Expected errors, got nil")
		}

		assert.Equal(t, "internal server error", err.Error())
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		m.AssertExpectations(t)
	})

	t.Run("logs error and returns bad gateway error", func(t *testing.T) {
		t.Helper()

		m := new(mock.ResultClient)
		logger, hook := test.NewNullLogger()
		client := g.NewResultClient(m, logger)

		stream := new(mock.ResultStream)

		request := proto.TeamResultRequest{
			TeamId: 1,
			Limit:  &wrappers.UInt64Value{Value: 8},
		}

		ctx := context.Background()

		e := status.Error(codes.Aborted, "aborted")

		m.On("GetResultsForTeam", ctx, &request, []grpc.CallOption(nil)).Return(stream, e)

		_, err := client.ByTeam(ctx, &request)

		if err == nil {
			t.Fatal("Expected errors, got nil")
		}

		assert.Equal(t, "error response returned from external service", err.Error())
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		m.AssertExpectations(t)
	})

	t.Run("logs error and returns internal server error if error parsing stream", func(t *testing.T) {
		t.Helper()

		m := new(mock.ResultClient)
		logger, hook := test.NewNullLogger()
		client := g.NewResultClient(m, logger)

		stream := new(mock.ResultStream)

		request := proto.TeamResultRequest{
			TeamId: 1,
			Limit:  &wrappers.UInt64Value{Value: 8},
		}

		ctx := context.Background()

		e := errors.New("oh damn")

		m.On("GetResultsForTeam", ctx, &request, []grpc.CallOption(nil)).Return(stream, nil)
		stream.On("Recv").Twice().Return(newProtoResult(), nil)
		stream.On("Recv").Once().Return(&proto.Result{}, e)

		_, err := client.ByTeam(ctx, &request)

		if err == nil {
			t.Fatal("Expected errors, got nil")
		}

		assert.Equal(t, "internal server error", err.Error())
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		m.AssertExpectations(t)
	})
}

func TestResultClient_ByID(t *testing.T) {
	t.Run("returns a result struct", func(t *testing.T) {
		t.Helper()

		m := new(mock.ResultClient)
		logger, _ := test.NewNullLogger()
		client := g.NewResultClient(m, logger)

		req := mock2.MatchedBy(func (r *proto.ResultRequest) bool {
			assert.Equal(t, uint64(78102), r.FixtureId)
			return true
		})

		ctx := context.Background()

		m.On("GetById", ctx, req, []grpc.CallOption(nil)).Return(newProtoResult(), nil)

		result, err := client.ByID(ctx, uint64(78102))

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		m.AssertExpectations(t)
		assertResult(t, result)
	})

	t.Run("returns not found error if returned by client", func(t *testing.T) {
		t.Helper()

		m := new(mock.ResultClient)
		logger, _ := test.NewNullLogger()
		client := g.NewResultClient(m, logger)

		req := mock2.MatchedBy(func (r *proto.ResultRequest) bool {
			assert.Equal(t, uint64(78102), r.FixtureId)
			return true
		})

		ctx := context.Background()

		e := status.Error(codes.NotFound, "not found")

		m.On("GetById", ctx, req, []grpc.CallOption(nil)).Return(&proto.Result{}, e)

		_, err := client.ByID(ctx, uint64(78102))

		if err == nil {
			t.Fatal("Expected error got nil")
		}

		a := assert.New(t)
		a.Equal(er.ErrorNotFound, err)
		m.AssertExpectations(t)
	})

	t.Run("logs error and returns internal server error if returned by client", func(t *testing.T) {
		t.Helper()

		m := new(mock.ResultClient)
		logger, hook := test.NewNullLogger()
		client := g.NewResultClient(m, logger)

		req := mock2.MatchedBy(func (r *proto.ResultRequest) bool {
			assert.Equal(t, uint64(78102), r.FixtureId)
			return true
		})

		ctx := context.Background()

		e := status.Error(codes.Internal, "internal server error")

		m.On("GetById", ctx, req, []grpc.CallOption(nil)).Return(&proto.Result{}, e)

		_, err := client.ByID(ctx, uint64(78102))

		if err == nil {
			t.Fatal("Expected error got nil")
		}

		a := assert.New(t)
		a.Equal(er.ErrorInternalServerError, err)
		assert.Equal(t, "internal server error", err.Error())
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		m.AssertExpectations(t)
	})

	t.Run("logs error and returns bad gateway error", func(t *testing.T) {
		t.Helper()

		m := new(mock.ResultClient)
		logger, hook := test.NewNullLogger()
		client := g.NewResultClient(m, logger)

		req := mock2.MatchedBy(func (r *proto.ResultRequest) bool {
			assert.Equal(t, uint64(78102), r.FixtureId)
			return true
		})

		ctx := context.Background()

		e := status.Error(codes.Aborted, "internal server error")

		m.On("GetById", ctx, req, []grpc.CallOption(nil)).Return(&proto.Result{}, e)

		_, err := client.ByID(ctx, uint64(78102))

		if err == nil {
			t.Fatal("Expected error got nil")
		}

		a := assert.New(t)
		a.Equal(er.ErrorBadGateway, err)
		assert.Equal(t, "error response returned from external service", err.Error())
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		m.AssertExpectations(t)
	})
}

func assertResult(t *testing.T, result *app.Result) {
	a := assert.New(t)

	homeCode := "WHU"
	homeFounded := uint64(1895)
	homeLogo := "logo"

	home := app.Team{
		ID:           1,
		Name:         "West Ham United",
		ShortCode:    &homeCode,
		CountryID:    8,
		VenueID:      214,
		NationalTeam: false,
		Founded:      &homeFounded,
		Logo:         &homeLogo,
	}

	awayCode := "NOT"
	awayFounded := uint64(1895)
	awayLogo := "logo"

	away := app.Team{
		ID:           10,
		Name:         "Nottingham Forest",
		ShortCode:    &awayCode,
		CountryID:    8,
		VenueID:      300,
		NationalTeam: true,
		Founded:      &awayFounded,
		Logo:         &awayLogo,
	}

	season := app.Season{
		ID:        16036,
		Name:      "2019/2020",
		IsCurrent: true,
	}

	start, _ := time.Parse(time.RFC3339, "2020-07-07T12:00:00+00:00")
	end, _ := time.Parse(time.RFC3339, "2020-07-23T23:59:59+00:00")

	round := app.Round{
		ID:        38,
		Name:      "38",
		SeasonID:  16036,
		StartDate: app.JsonDate(start),
		EndDate:   app.JsonDate(end),
	}

	venue := app.Venue{
		ID:   214,
		Name: "London Stadium",
	}

	stats := app.ResultStats{
		HomeScore: 5,
		AwayScore: 2,
	}

	date, _ := time.Parse(time.RFC3339, "2020-07-07T15:00:00+00:00")

	resultDate := app.JsonDate(date)

	a.Equal(uint64(78102), result.ID)
	a.Equal(home, result.HomeTeam)
	a.Equal(away, result.AwayTeam)
	a.Equal(season, result.Season)
	a.Equal(round, *result.Round)
	a.Equal(venue, result.Venue)
	a.Equal(stats, result.Stats)
	a.Equal(resultDate, result.DateTime)
}

func newProtoResult() *proto.Result {
	home := proto.Team{
		Id:             1,
		Name:           "West Ham United",
		ShortCode:      &wrappers.StringValue{Value: "WHU"},
		CountryId:      8,
		VenueId:        214,
		IsNationalTeam: &wrappers.BoolValue{Value: false},
		Founded:        &wrappers.UInt64Value{Value: 1895},
		Logo:           &wrappers.StringValue{Value: "logo"},
	}

	away := proto.Team{
		Id:             10,
		Name:           "Nottingham Forest",
		ShortCode:      &wrappers.StringValue{Value: "NOT"},
		CountryId:      8,
		VenueId:        300,
		IsNationalTeam: &wrappers.BoolValue{Value: true},
		Founded:        &wrappers.UInt64Value{Value: 1895},
		Logo:           &wrappers.StringValue{Value: "logo"},
	}

	season := proto.Season{
		Id:        16036,
		Name:      "2019/2020",
		IsCurrent: &wrappers.BoolValue{Value: true},
	}

	round := proto.Round{
		Id:        38,
		Name:      "38",
		SeasonId:  16036,
		StartDate: "2020-07-07T12:00:00+00:00",
		EndDate:   "2020-07-23T23:59:59+00:00",
	}

	venue := proto.Venue{
		Id:   214,
		Name: "London Stadium",
	}

	date := proto.Date{
		Utc: 1594132077,
		Rfc: "2020-07-07T15:00:00+00:00",
	}

	stats := proto.MatchStats{
		HomeScore: &wrappers.UInt32Value{Value: 5},
		AwayScore: &wrappers.UInt32Value{Value: 2},
	}

	return &proto.Result{
		Id:       78102,
		HomeTeam: &home,
		AwayTeam: &away,
		Season:   &season,
		Round:    &round,
		Venue:    &venue,
		DateTime: &date,
		Stats:    &stats,
	}
}
