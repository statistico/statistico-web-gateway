package composer_test

import (
	"context"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/composer"
	"github.com/statistico/statistico-web-gateway/internal/app/errors"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc/proto"
	"github.com/statistico/statistico-web-gateway/internal/app/mock"
	"github.com/stretchr/testify/assert"
	mock2 "github.com/stretchr/testify/mock"
	"testing"
)

func TestTeamStatComposer_FetchStats(t *testing.T) {
	t.Run("calls team stat client and returns a slice of team stat struct", func (t *testing.T) {
		t.Helper()

		client := new(mock.TeamStatClient)
		resultClient := new(mock.ResultGRPCClient)
		comp := composer.NewTeamStatComposer(client, resultClient)

		team := composer.TeamFilter{ID: 10}
		seasonIds := []uint64{16036}

		filters := composer.TeamStatFilters{
			SeasonIds:  &seasonIds,
			Stat:       "goals",
			Team:       team,
		}

		stats := make([]*app.TeamStat, 2)
		stats[0] = newTeamStat(1, 5)
		stats[1] = newTeamStat(2, 3)

		req := mock2.MatchedBy(func (r *proto.TeamStatRequest) bool {
			assert.Equal(t, uint64(10), r.TeamId)
			assert.Equal(t, "goals", r.Stat)
			assert.Equal(t, []uint64{16036}, r.SeasonIds)
			return true
		})

		ch := teamStatChannel(stats)
		errChan := make(chan error)
		close(errChan)

		client.On("Stats", context.Background(), req).Return(ch, errChan)

		ctx := context.Background()

		fetched, err := comp.FetchStats(ctx, &filters)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Equal(t, stats, fetched)
		client.AssertExpectations(t)
	})

	t.Run("returns an error if returned by grpc client", func(t *testing.T) {
		t.Helper()

		client := new(mock.TeamStatClient)
		resultClient := new(mock.ResultGRPCClient)
		comp := composer.NewTeamStatComposer(client, resultClient)

		team := composer.TeamFilter{ID: 10}
		seasonIds := []uint64{16036}

		filters := composer.TeamStatFilters{
			SeasonIds:  &seasonIds,
			Stat:       "goals",
			Team:       team,
		}

		req := mock2.MatchedBy(func (r *proto.TeamStatRequest) bool {
			assert.Equal(t, uint64(10), r.TeamId)
			assert.Equal(t, "goals", r.Stat)
			assert.Equal(t, []uint64{16036}, r.SeasonIds)
			return true
		})

		stats := make([]*app.TeamStat, 2)
		stats[0] = newTeamStat(1, 5)
		stats[1] = newTeamStat(2, 3)

		statChan := teamStatChannel(stats)
		errChan := errorChannel(errors.ErrorBadGateway)

		client.On("Stats", context.Background(), req).Return(statChan, errChan)

		ctx := context.Background()

		_, err := comp.FetchStats(ctx, &filters)

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, errors.ErrorBadGateway, err)
		client.AssertExpectations(t)
	})

	t.Run("calls result client to fetch associated result if includes result parameter provided", func(t *testing.T) {
		t.Helper()

		client := new(mock.TeamStatClient)
		resultClient := new(mock.ResultGRPCClient)
		comp := composer.NewTeamStatComposer(client, resultClient)

		team := composer.TeamFilter{ID: 10}
		seasonIds := []uint64{16036}
		include := []string{"result"}

		filters := composer.TeamStatFilters{
			Include: 	&include,
			SeasonIds:  &seasonIds,
			Stat:       "goals",
			Team:       team,
		}

		stats := make([]*app.TeamStat, 1)
		stats[0] = newTeamStat(1, 5)

		req := mock2.MatchedBy(func (r *proto.TeamStatRequest) bool {
			assert.Equal(t, uint64(10), r.TeamId)
			assert.Equal(t, "goals", r.Stat)
			assert.Equal(t, []uint64{16036}, r.SeasonIds)
			return true
		})

		result := newResult(1)

		ctx := context.Background()
		ch := teamStatChannel(stats)
		errChan := make(chan error)
		close(errChan)

		client.On("Stats", ctx, req).Return(ch, errChan)
		resultClient.On("ByID", ctx, uint64(1)).Return(result, nil)

		fetched, err := comp.FetchStats(ctx, &filters)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Equal(t, stats, fetched)
		assert.Equal(t, result, fetched[0].Result)
		client.AssertExpectations(t)
		resultClient.AssertExpectations(t)
	})

	t.Run("returns error if returned by result client", func(t *testing.T) {
		t.Helper()

		client := new(mock.TeamStatClient)
		resultClient := new(mock.ResultGRPCClient)
		comp := composer.NewTeamStatComposer(client, resultClient)

		team := composer.TeamFilter{ID: 10}
		seasonIds := []uint64{16036}
		include := []string{"result"}

		filters := composer.TeamStatFilters{
			Include: 	&include,
			SeasonIds:  &seasonIds,
			Stat:       "goals",
			Team:       team,
		}

		stats := make([]*app.TeamStat, 1)
		stats[0] = newTeamStat(1, 5)

		req := mock2.MatchedBy(func (r *proto.TeamStatRequest) bool {
			assert.Equal(t, uint64(10), r.TeamId)
			assert.Equal(t, "goals", r.Stat)
			assert.Equal(t, []uint64{16036}, r.SeasonIds)
			return true
		})

		result := newResult(1)

		ctx := context.Background()
		ch := teamStatChannel(stats)
		errChan := make(chan error, 1)

		client.On("Stats", ctx, req).Return(ch, errChan)
		resultClient.On("ByID", ctx, uint64(1)).Return(result, errors.ErrorNotFound)

		_, err := comp.FetchStats(ctx, &filters)

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, err, errors.ErrorNotFound)
		client.AssertExpectations(t)
		resultClient.AssertExpectations(t)
	})
}

func newTeamStat(fixtureID uint64, value uint32) *app.TeamStat {
	return &app.TeamStat{
		FixtureID: fixtureID,
		Stat:      "goals",
		Value:     &value,
	}
}

func newResult(fixtureID uint64) *app.Result {
	return &app.Result{ID: fixtureID}
}

func teamStatChannel(stats []*app.TeamStat) <-chan *app.TeamStat {
	ch := make(chan *app.TeamStat, len(stats))

	for _, s := range stats {
		ch <- s
	}

	close(ch)

	return ch
}

func errorChannel(err error) chan error {
	ch := make(chan error, 1)
	ch <- err
	close(ch)
	return ch
}
