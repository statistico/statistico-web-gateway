package composer

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc/proto"
	"time"
)

type ResultComposer interface {
	ForTeam(teamId uint64, filters *Filters) ([]*app.Result, error)
}

type resultComposer struct {
	client grpc.ResultClient
}

func (r resultComposer) ForTeam(teamId uint64, filters *Filters) ([]*app.Result, error) {
	request := proto.TeamResultRequest{TeamId: teamId}

	if filters.Limit != nil {
		request.Limit = &wrappers.UInt64Value{Value: *filters.Limit}
	}

	if filters.DateBefore != nil {
		request.DateBefore = &wrappers.StringValue{Value: filters.DateBefore.Format(time.RFC3339)}
	}

	if filters.DateAfter != nil {
		request.DateAfter = &wrappers.StringValue{Value: filters.DateAfter.Format(time.RFC3339)}
	}

	if filters.Venue != nil {
		request.Venue = &wrappers.StringValue{Value: *filters.Venue}
	}

	if filters.Sort != nil {
		request.Sort = &wrappers.StringValue{Value: *filters.Sort}
	}

	return r.client.ByTeam(context.Background(), &request)
}

func NewResultComposer(c grpc.ResultClient) ResultComposer {
	return &resultComposer{client: c}
}
