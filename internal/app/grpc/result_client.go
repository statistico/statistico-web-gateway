package grpc

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/errors"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

type ResultClient interface {
	ByID(ctx context.Context, fixtureID uint64) (*app.Result, error)
	ByTeam(ctx context.Context, req *proto.TeamResultRequest) ([]*app.Result, error)
}

type resultClient struct {
	client proto.ResultServiceClient
	logger *logrus.Logger
}

func (r resultClient) ByID(ctx context.Context, fixtureID uint64) (*app.Result, error) {
	request := proto.ResultRequest{FixtureId: fixtureID}

	result, err := r.client.GetById(ctx, &request)

	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				return nil, errors.ErrorNotFound
			case codes.Internal:
				r.logError(err)
				return nil, errors.ErrorInternalServerError
			default:
				r.logError(err)
				return nil, errors.ErrorBadGateway
			}
		}
	}

	converted, err := convertResult(result)

	if err != nil {
		r.logError(err)
		return nil, errors.ErrorInternalServerError
	}

	return converted, nil
}

func (r resultClient) ByTeam(ctx context.Context, req *proto.TeamResultRequest) ([]*app.Result, error) {
	results := []*app.Result{}

	stream, err := r.client.GetResultsForTeam(ctx, req)

	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.InvalidArgument:
				return results, err
			case codes.Internal:
				r.logError(err)
				return results, errors.ErrorInternalServerError
			default:
				r.logError(err)
				return results, errors.ErrorBadGateway
			}
		}
	}

	for {
		result, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			r.logError(err)
			return results, errors.ErrorInternalServerError
		}

		res, err := convertResult(result)

		if err != nil {
			r.logError(err)
			return results, errors.ErrorInternalServerError
		}

		results = append(results, res)
	}

	return results, nil
}

func (r resultClient) logError(err error) {
	r.logger.Errorf("Error in result client %s", err.Error())
}

func NewResultClient(p proto.ResultServiceClient, l *logrus.Logger) ResultClient {
	return resultClient{client: p, logger: l}
}
