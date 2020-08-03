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

func TestSeasonClient_ByTeamId(t *testing.T) {
	t.Run("calls season client and returns a slice of season struct", func(t *testing.T) {
		t.Helper()

		s := new(mock.SeasonClient)
		logger, hook := test.NewNullLogger()
		client := g.NewSeasonClient(s, logger)

		stream := new(mock.SeasonStream)

		request := proto.TeamSeasonsRequest{
			TeamId: 55,
			Sort:   &wrappers.StringValue{Value: "name_desc"},
		}

		ctx := context.Background()

		s.On("GetSeasonsForTeam", ctx, &request, []grpc.CallOption(nil)).Return(stream, nil)
		stream.On("Recv").Twice().Return(newProtoSeason(), nil)
		stream.On("Recv").Once().Return(&proto.Season{}, io.EOF)

		seasons, err := client.ByTeamId(ctx, 55, "name_desc")

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

		s := new(mock.SeasonClient)
		logger, hook := test.NewNullLogger()
		client := g.NewSeasonClient(s, logger)

		stream := new(mock.SeasonStream)

		request := proto.TeamSeasonsRequest{
			TeamId: 55,
			Sort:   &wrappers.StringValue{Value: "name_desc"},
		}

		ctx := context.Background()

		e := status.Error(codes.Internal, "internal error")

		s.On("GetSeasonsForTeam", ctx, &request, []grpc.CallOption(nil)).Return(stream, e)

		_, err := client.ByTeamId(ctx, 55, "name_desc")

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

		s := new(mock.SeasonClient)
		logger, hook := test.NewNullLogger()
		client := g.NewSeasonClient(s, logger)

		stream := new(mock.SeasonStream)

		request := proto.TeamSeasonsRequest{
			TeamId: 55,
			Sort:   &wrappers.StringValue{Value: "name_desc"},
		}

		ctx := context.Background()

		e := status.Error(codes.Unavailable, "service unavailable")

		s.On("GetSeasonsForTeam", ctx, &request, []grpc.CallOption(nil)).Return(stream, e)

		_, err := client.ByTeamId(ctx, 55, "name_desc")

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

		s := new(mock.SeasonClient)
		logger, hook := test.NewNullLogger()
		client := g.NewSeasonClient(s, logger)

		stream := new(mock.SeasonStream)

		request := proto.TeamSeasonsRequest{
			TeamId: 55,
			Sort:   &wrappers.StringValue{Value: "name_desc"},
		}

		ctx := context.Background()

		e := errors.New("oh damn")

		s.On("GetSeasonsForTeam", ctx, &request, []grpc.CallOption(nil)).Return(stream, nil)
		stream.On("Recv").Twice().Return(newProtoSeason(), nil)
		stream.On("Recv").Once().Return(&proto.Season{}, e)

		_, err := client.ByTeamId(ctx, 55, "name_desc")

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "internal server error", err.Error())
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		assert.Equal(t, "Error in competition client: oh damn", hook.LastEntry().Message)
		s.AssertExpectations(t)
		stream.AssertExpectations(t)
	})
}
