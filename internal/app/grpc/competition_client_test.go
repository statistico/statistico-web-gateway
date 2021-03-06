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

func TestCompetitionClient_CompetitionByCountryId(t *testing.T) {
	t.Run("calls competition client and returns a slice of competition struct", func(t *testing.T) {
		t.Helper()

		m := new(mock.CompetitionClient)
		s := new(mock.SeasonClient)
		logger, hook := test.NewNullLogger()
		client := g.NewCompetitionClient(m, s, logger)

		stream := new(mock.CompetitionStream)

		request := proto.CompetitionRequest{
			CountryIds: []uint64{462},
			Sort:       nil,
			IsCup:      nil,
		}

		ctx := context.Background()

		m.On("ListCompetitions", ctx, &request, []grpc.CallOption(nil)).Return(stream, nil)
		stream.On("Recv").Twice().Return(newProtoCompetition(), nil)
		stream.On("Recv").Once().Return(&proto.Competition{}, io.EOF)

		competitions, err := client.CompetitionsByCountryId(ctx, 462)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Equal(t, 2, len(competitions))
		assert.Nil(t, hook.LastEntry())
		m.AssertExpectations(t)
		stream.AssertExpectations(t)
	})

	t.Run("logs error and returns internal server error if internal server error returned by client", func(t *testing.T) {
		t.Helper()

		m := new(mock.CompetitionClient)
		s := new(mock.SeasonClient)
		logger, hook := test.NewNullLogger()
		client := g.NewCompetitionClient(m, s, logger)

		stream := new(mock.CompetitionStream)

		request := proto.CompetitionRequest{
			CountryIds: []uint64{462},
			Sort:       nil,
			IsCup:      nil,
		}

		ctx := context.Background()

		e := status.Error(codes.Internal, "internal error")

		m.On("ListCompetitions", ctx, &request, []grpc.CallOption(nil)).Return(stream, e)

		_, err := client.CompetitionsByCountryId(ctx, 462)

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "internal server error", err.Error())
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		assert.Equal(t, "Error in competition client: rpc error: code = Internal desc = internal error", hook.LastEntry().Message)
		m.AssertExpectations(t)
		stream.AssertNotCalled(t, "Recv")
	})

	t.Run("logs error and returns bad gateway error for non internal server error returned by client", func(t *testing.T) {
		t.Helper()

		m := new(mock.CompetitionClient)
		s := new(mock.SeasonClient)
		logger, hook := test.NewNullLogger()
		client := g.NewCompetitionClient(m, s, logger)

		stream := new(mock.CompetitionStream)

		request := proto.CompetitionRequest{
			CountryIds: []uint64{462},
			Sort:       nil,
			IsCup:      nil,
		}

		ctx := context.Background()

		e := status.Error(codes.Unavailable, "service unavailable")

		m.On("ListCompetitions", ctx, &request, []grpc.CallOption(nil)).Return(stream, e)

		_, err := client.CompetitionsByCountryId(ctx, 462)

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "error response returned from external service", err.Error())
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		assert.Equal(t, "Error in competition client: rpc error: code = Unavailable desc = service unavailable", hook.LastEntry().Message)
		m.AssertExpectations(t)
		stream.AssertNotCalled(t, "Recv")
	})

	t.Run("logs error and returns internal server error if error reading from stream", func(t *testing.T) {
		t.Helper()

		m := new(mock.CompetitionClient)
		s := new(mock.SeasonClient)
		logger, hook := test.NewNullLogger()
		client := g.NewCompetitionClient(m, s, logger)

		stream := new(mock.CompetitionStream)

		request := proto.CompetitionRequest{
			CountryIds: []uint64{462},
			Sort:       nil,
			IsCup:      nil,
		}

		ctx := context.Background()

		e := errors.New("oh damn")

		m.On("ListCompetitions", ctx, &request, []grpc.CallOption(nil)).Return(stream, nil)
		stream.On("Recv").Twice().Return(newProtoCompetition(), nil)
		stream.On("Recv").Once().Return(&proto.Competition{}, e)

		_, err := client.CompetitionsByCountryId(ctx, 462)

		if err == nil {
			t.Fatal("Expected errors, got nil")
		}

		assert.Equal(t, "internal server error", err.Error())
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		assert.Equal(t, "Error in competition client: oh damn", hook.LastEntry().Message)
		m.AssertExpectations(t)
		stream.AssertExpectations(t)
	})
}

func TestCompetitionClient_CompetitionSeasons(t *testing.T) {
	t.Run("calls season client and returns a slice of season struct", func(t *testing.T) {
		t.Helper()

		m := new(mock.CompetitionClient)
		s := new(mock.SeasonClient)
		logger, hook := test.NewNullLogger()
		client := g.NewCompetitionClient(m, s, logger)

		stream := new(mock.SeasonStream)

		request := proto.SeasonCompetitionRequest{
			CompetitionId: 8,
			Sort:          &wrappers.StringValue{Value: "name_asc"},
		}

		ctx := context.Background()

		s.On("GetSeasonsForCompetition", ctx, &request, []grpc.CallOption(nil)).Return(stream, nil)
		stream.On("Recv").Twice().Return(newProtoSeason(), nil)
		stream.On("Recv").Once().Return(&proto.Season{}, io.EOF)

		seasons, err := client.CompetitionSeasons(ctx, 8, "name_asc")

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Equal(t, 2, len(seasons))
		assert.Nil(t, hook.LastEntry())
		s.AssertExpectations(t)
		stream.AssertExpectations(t)
	})

	t.Run("logs error and returns internal server error if internal server error returned by client", func(t *testing.T) {
		t.Helper()

		m := new(mock.CompetitionClient)
		s := new(mock.SeasonClient)
		logger, hook := test.NewNullLogger()
		client := g.NewCompetitionClient(m, s, logger)

		stream := new(mock.SeasonStream)

		request := proto.SeasonCompetitionRequest{
			CompetitionId: 8,
			Sort:          &wrappers.StringValue{Value: "name_asc"},
		}

		ctx := context.Background()

		e := status.Error(codes.Internal, "internal error")

		s.On("GetSeasonsForCompetition", ctx, &request, []grpc.CallOption(nil)).Return(stream, e)

		_, err := client.CompetitionSeasons(ctx, 8, "name_asc")

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "internal server error", err.Error())
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		assert.Equal(t, "Error in competition client: rpc error: code = Internal desc = internal error", hook.LastEntry().Message)
		s.AssertExpectations(t)
		stream.AssertNotCalled(t, "Recv")
	})

	t.Run("logs error and returns bad gateway error for non internal server error returned by client", func(t *testing.T) {
		t.Helper()

		m := new(mock.CompetitionClient)
		s := new(mock.SeasonClient)
		logger, hook := test.NewNullLogger()
		client := g.NewCompetitionClient(m, s, logger)

		stream := new(mock.SeasonStream)

		request := proto.SeasonCompetitionRequest{
			CompetitionId: 8,
			Sort:          &wrappers.StringValue{Value: "name_asc"},
		}

		ctx := context.Background()

		e := status.Error(codes.Unavailable, "service unavailable")

		s.On("GetSeasonsForCompetition", ctx, &request, []grpc.CallOption(nil)).Return(stream, e)

		_, err := client.CompetitionSeasons(ctx, 8, "name_asc")

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "error response returned from external service", err.Error())
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		assert.Equal(t, "Error in competition client: rpc error: code = Unavailable desc = service unavailable", hook.LastEntry().Message)
		s.AssertExpectations(t)
		stream.AssertNotCalled(t, "Recv")
	})

	t.Run("logs error and returns internal server error if error reading from stream", func(t *testing.T) {
		t.Helper()

		m := new(mock.CompetitionClient)
		s := new(mock.SeasonClient)
		logger, hook := test.NewNullLogger()
		client := g.NewCompetitionClient(m, s, logger)

		stream := new(mock.SeasonStream)

		request := proto.SeasonCompetitionRequest{
			CompetitionId: 8,
			Sort:          &wrappers.StringValue{Value: "name_asc"},
		}

		ctx := context.Background()

		e := errors.New("oh damn")

		s.On("GetSeasonsForCompetition", ctx, &request, []grpc.CallOption(nil)).Return(stream, nil)
		stream.On("Recv").Twice().Return(newProtoSeason(), nil)
		stream.On("Recv").Once().Return(&proto.Season{}, e)

		_, err := client.CompetitionSeasons(ctx, 8, "name_asc")

		if err == nil {
			t.Fatal("Expected errors, got nil")
		}

		assert.Equal(t, "internal server error", err.Error())
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		assert.Equal(t, "Error in competition client: oh damn", hook.LastEntry().Message)
		s.AssertExpectations(t)
		stream.AssertExpectations(t)
	})
}

func newProtoCompetition() *proto.Competition {
	return &proto.Competition{
		Id:        8,
		Name:      "Premier League",
		IsCup:     false,
		CountryId: 462,
	}
}

func newProtoSeason() *proto.Season {
	return &proto.Season{
		Id:        16036,
		Name:      "2019/2020",
		IsCurrent: &wrappers.BoolValue{Value: true},
	}
}
