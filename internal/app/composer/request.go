package composer

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc/proto"
	"time"
)

func teamStatRequestFromFilters(filters *TeamStatFilters) *proto.TeamStatRequest {
	request := proto.TeamStatRequest{TeamId: filters.Team.ID, Stat: filters.Stat}

	if filters.DateAfter != nil {
		request.DateAfter = &wrappers.StringValue{Value: filters.DateAfter.Format(time.RFC3339)}
	}

	if filters.DateBefore != nil {
		request.DateBefore = &wrappers.StringValue{Value: filters.DateBefore.Format(time.RFC3339)}
	}

	if filters.Limit != nil {
		request.Limit = &wrappers.UInt64Value{Value: *filters.Limit}
	}

	if filters.Opponent != nil {
		request.Opponent = &wrappers.BoolValue{Value: *filters.Opponent}
	}

	if filters.SeasonIds != nil {
		request.SeasonIds = *filters.SeasonIds
	}

	if filters.Sort != nil {
		request.Sort = &wrappers.StringValue{Value: *filters.Sort}
	}

	if filters.Team.Venue != nil {
		request.Venue = &wrappers.StringValue{Value: *filters.Team.Venue}
	}

	return &request
}

