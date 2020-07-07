package grpc

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc/proto"
	"github.com/statistico/statistico-web-gateway/internal/app/mock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"testing"
)

func TestResultClient_ByTeam(t *testing.T) {
	t.Run("calls result client and returns a slice of result struct", func(t *testing.T) {
		t.Helper()

		m := new(mock.ResultClient)
		logger, hook := test.NewNullLogger()
		client := NewResultClient(m, logger)

		stream := new(mock.ResultStream)

		request := proto.TeamResultRequest{
			TeamId:     1,
			Limit:      &wrappers.UInt64Value{Value: 8},
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
		client := NewResultClient(m, logger)

		stream := new(mock.ResultStream)

		request := proto.TeamResultRequest{
			TeamId:     1,
			Limit:      &wrappers.UInt64Value{Value: 8},
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
		client := NewResultClient(m, logger)

		stream := new(mock.ResultStream)

		request := proto.TeamResultRequest{
			TeamId:     1,
			Limit:      &wrappers.UInt64Value{Value: 8},
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
		client := NewResultClient(m, logger)

		stream := new(mock.ResultStream)

		request := proto.TeamResultRequest{
			TeamId:     1,
			Limit:      &wrappers.UInt64Value{Value: 8},
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
		client := NewResultClient(m, logger)

		stream := new(mock.ResultStream)

		request := proto.TeamResultRequest{
			TeamId:     1,
			Limit:      &wrappers.UInt64Value{Value: 8},
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
