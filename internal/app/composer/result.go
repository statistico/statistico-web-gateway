package composer

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/statistico/statistico-data-go-grpc-client"
	"github.com/statistico/statistico-proto/go"
	"time"
)

type ResultComposer interface {
	FetchResults(filters *ResultFilters) ([]*statisticoproto.Result, error)
}

type resultComposer struct {
	client statisticodata.ResultClient
}

func (r *resultComposer) FetchResults(filters *ResultFilters) ([]*statisticoproto.Result, error) {
	if filters.Team != nil {
		request := statisticoproto.TeamResultRequest{TeamId: filters.Team.ID}

		if filters.Limit != nil {
			request.Limit = &wrappers.UInt64Value{Value: *filters.Limit}
		}

		if filters.DateBefore != nil {
			request.DateBefore = &wrappers.StringValue{Value: filters.DateBefore.Format(time.RFC3339)}
		}

		if filters.DateAfter != nil {
			request.DateAfter = &wrappers.StringValue{Value: filters.DateAfter.Format(time.RFC3339)}
		}

		if filters.Team.Venue != nil {
			request.Venue = &wrappers.StringValue{Value: *filters.Team.Venue}
		}

		if filters.SeasonIds != nil {
			request.SeasonIds = *filters.SeasonIds
		}

		if filters.Sort != nil {
			request.Sort = &wrappers.StringValue{Value: *filters.Sort}
		}

		return r.client.ByTeam(context.Background(), &request)
	}

	return []*statisticoproto.Result{}, nil
}

func NewResultComposer(c statisticodata.ResultClient) ResultComposer {
	return &resultComposer{client: c}
}
