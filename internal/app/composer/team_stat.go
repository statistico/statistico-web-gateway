package composer

import (
	"context"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc"
)

type TeamStatComposer interface {
	FetchStats(ctx context.Context, filters *TeamStatFilters) ([]*app.TeamStat, error)
}

type teamStatComposer struct {
	statClient grpc.TeamStatClient
	resultClient grpc.ResultClient
}

func (t *teamStatComposer) FetchStats(ctx context.Context, filters *TeamStatFilters) ([]*app.TeamStat, error) {
	request := teamStatRequestFromFilters(filters)

	statChan, errChan := t.statClient.Stats(ctx, request)

	stats := []*app.TeamStat{}

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
	return &teamStatComposer{statClient: c, resultClient: r}
}
