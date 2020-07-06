package grpc

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-web-gateway/internal/app/errors"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ResultClient interface {
	ByTeam(ctx context.Context, req *proto.TeamResultRequest) ([]proto.Result, error)
}

type resultClient struct {
	client proto.ResultServiceClient
	logger *logrus.Logger
}

func (r resultClient) ByTeam(ctx context.Context, req *proto.TeamResultRequest) ([]proto.Result, error) {
	response, err := r.client.GetResultsForTeam(ctx, req)

	var res []proto.Result

	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.InvalidArgument:
				return res, err
			case codes.Internal:
				r.logError(err)
				return res, errors.ErrorInternalServerError
			default:
				r.logError(err)
				return res, errors.ErrorBadGateway
			}

		}
	}

	for _, result := range response {

	}

	return res, nil
}

func (r resultClient) logError(err error) {
	r.logger.Errorf("Error in result client %s", err.Error())
}