package composer

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc/proto"
	"time"
)

type TeamStatComposer interface {
	FetchStats(ctx context.Context, filters *TeamStatFilters) ([]*app.TeamStat, error)
}

type teamStatComposer struct {
	client grpc.TeamStatClient
	resultClient grpc.ResultClient
}

func (t *teamStatComposer) FetchStats(ctx context.Context, filters *TeamStatFilters) ([]*app.TeamStat, error) {
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

	stats := []*app.TeamStat{}

	statChan, errChan := t.client.Stats(ctx, &request)

	for stat := range statChan {
		if filters.IncludesParameter("result") {
			stat.Result = t.parseAssociatedResult(ctx, stat, errChan)
		}

		stats = append(stats, stat)
	}

	for err := range errChan {
		return stats, err
	}

	return stats, nil
}

func (t *teamStatComposer) parseAssociatedResult(ctx context.Context, stat *app.TeamStat, errChan chan error) *app.Result {
	result, err := t.resultClient.ByID(ctx, stat.FixtureID)

	if err != nil {
		errChan <- err
		return nil
	}

	return result
}

func NewTeamStatComposer(c grpc.TeamStatClient, r grpc.ResultClient) TeamStatComposer {
	return &teamStatComposer{client: c, resultClient: r}
}
