package grpc_test

import (
	"context"
	"errors"
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

func TestTeamStatClient_Stats(t *testing.T) {
	t.Run("calls team stat client and returns a slice of team stat struct", func(t *testing.T) {
		t.Helper()

		m := new(mock.TeamStatClient)
		logger, hook := test.NewNullLogger()
		client := g.NewTeamStatClient(m, logger)

		stream := new(mock.TeamStatStream)

		request := proto.TeamStatRequest{
			Stat:       "shots_total",
			TeamId:     5,
			SeasonIds:  []uint64{16036},
		}

		ctx := context.Background()

		m.On("GetStatForTeam", ctx, &request, []grpc.CallOption(nil)).Return(stream, nil)
		stream.On("Recv").Twice().Return(newProtoTeamStat(), nil)
		stream.On("Recv").Once().Return(&proto.TeamStat{}, io.EOF)

		stats, err := client.Stats(ctx, &request)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Equal(t, 2, len(stats))
		assert.Nil(t, hook.LastEntry())
		m.AssertExpectations(t)
	})

	t.Run("returns error if invalid argument error returned by team stat client", func(t *testing.T) {
		t.Helper()

		m := new(mock.TeamStatClient)
		logger, hook := test.NewNullLogger()
		client := g.NewTeamStatClient(m, logger)

		stream := new(mock.TeamStatStream)

		request := proto.TeamStatRequest{
			Stat:       "shots_total",
			TeamId:     5,
			SeasonIds:  []uint64{16036},
		}

		ctx := context.Background()

		e := status.Error(codes.InvalidArgument, "incorrect format")

		m.On("GetStatForTeam", ctx, &request, []grpc.CallOption(nil)).Return(stream, e)

		_, err := client.Stats(ctx, &request)

		if err == nil {
			t.Fatal("Expected errors, got nil")
		}

		assert.Equal(t, "rpc error: code = InvalidArgument desc = incorrect format", err.Error())
		assert.Nil(t, hook.LastEntry())
		m.AssertExpectations(t)
	})

	t.Run("logs error and returns internal server error", func(t *testing.T) {
		t.Helper()

		m := new(mock.TeamStatClient)
		logger, hook := test.NewNullLogger()
		client := g.NewTeamStatClient(m, logger)

		stream := new(mock.TeamStatStream)

		request := proto.TeamStatRequest{
			Stat:       "shots_total",
			TeamId:     5,
			SeasonIds:  []uint64{16036},
		}

		ctx := context.Background()

		e := status.Error(codes.Internal, "internal error")

		m.On("GetStatForTeam", ctx, &request, []grpc.CallOption(nil)).Return(stream, e)

		_, err := client.Stats(ctx, &request)

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

		m := new(mock.TeamStatClient)
		logger, hook := test.NewNullLogger()
		client := g.NewTeamStatClient(m, logger)

		stream := new(mock.TeamStatStream)

		request := proto.TeamStatRequest{
			Stat:       "shots_total",
			TeamId:     5,
			SeasonIds:  []uint64{16036},
		}

		ctx := context.Background()

		e := status.Error(codes.Aborted, "aborted")

		m.On("GetStatForTeam", ctx, &request, []grpc.CallOption(nil)).Return(stream, e)

		_, err := client.Stats(ctx, &request)

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

		m := new(mock.TeamStatClient)
		logger, hook := test.NewNullLogger()
		client := g.NewTeamStatClient(m, logger)

		stream := new(mock.TeamStatStream)

		request := proto.TeamStatRequest{
			Stat:       "shots_total",
			TeamId:     5,
			SeasonIds:  []uint64{16036},
		}

		ctx := context.Background()

		e := errors.New("oh damn")

		m.On("GetStatForTeam", ctx, &request, []grpc.CallOption(nil)).Return(stream, nil)
		stream.On("Recv").Twice().Return(newProtoTeamStat(), nil)
		stream.On("Recv").Once().Return(&proto.TeamStat{}, e)

		_, err := client.Stats(ctx, &request)

		if err == nil {
			t.Fatal("Expected errors, got nil")
		}

		assert.Equal(t, "internal server error", err.Error())
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		m.AssertExpectations(t)
	})
}

func newProtoTeamStat() *proto.TeamStat {
	return &proto.TeamStat{FixtureId: 43, Stat: "shots_total"}
}