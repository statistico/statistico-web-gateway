package composer_test

import (
	"context"
	"errors"
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/composer"
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
		comp := composer.NewTeamStatComposer(client)

		team := composer.TeamFilter{ID: 10}
		seasonIds := []uint64{16036}

		filters := composer.TeamStatFilters{
			SeasonIds:  &seasonIds,
			Stat:       "goals",
			Team:       team,
		}

		goals := uint32(3)

		stat := app.TeamStat{
			FixtureID: 43,
			Stat:      "goals",
			Value:     &goals,
		}

		stats := []*app.TeamStat{&stat}

		req := mock2.MatchedBy(func (r *proto.TeamStatRequest) bool {
			assert.Equal(t, uint64(10), r.TeamId)
			assert.Equal(t, "goals", r.Stat)
			assert.Equal(t, []uint64{16036}, r.SeasonIds)
			return true
		})

		client.On("Stats", context.Background(), req).Return(stats, nil)

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
		comp := composer.NewTeamStatComposer(client)

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

		returnedError := errors.New("oh no")

		client.On("Stats", context.Background(), req).Return([]*app.TeamStat{}, returnedError)

		ctx := context.Background()

		_, err := comp.FetchStats(ctx, &filters)

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, returnedError, err)
		client.AssertExpectations(t)
	})
}
