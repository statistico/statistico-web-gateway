package composer

import (
	"github.com/statistico/statistico-web-gateway/internal/app"
	"github.com/statistico/statistico-web-gateway/internal/app/proxy"
	"time"
)

type FixtureSearchQuery struct {
	DateFrom          time.Time         `json:"dateTo"`
	DateTo            time.Time         `json:"dateFrom"`
	SeasonIds         []uint64          `json:"seasonId"`
	StatFilter        app.StatFilter    `json:"statFilter"`
}

type FixtureHandler interface {
	Search(q *FixtureSearchQuery) (map[uint64][]*app.Fixture, error)
}

type FixtureSearch struct{
	performance *proxy.PerformanceService
	fixture     *proxy.FixtureService
}

func (f FixtureSearch) Search(q *FixtureSearchQuery) (map[uint64][]*app.Fixture, error) {
	fixtures := map[uint64][]*app.Fixture{}

	for _, id := range q.SeasonIds {
		fetched, err := f.fixture.FixtureForSeasonBetween(id, q.DateFrom, q.DateTo)

		if err != nil {
			return nil, err
		}

		teams, err := f.performance.ProxyTeamsMatchingFilterRequest(q.StatFilter, []uint64{id})

		if err != nil {
			return nil, err
		}

		fixtures[id] = filterFixtures(q.StatFilter.Venue, fetched, teams)
	}

	return fixtures, nil
}

func filterFixtures(venue string, fix []*app.Fixture, teams []*app.Team) []*app.Fixture {
	var fixtures []*app.Fixture

	for _, fix := range fix {
		for _, team := range teams {
			if venue == "home" {
				if fix.HomeTeam.ID == team.ID {
					fixtures = append(fixtures, fix)
				}
			}

			if venue == "away" {
				if fix.AwayTeam.ID == team.ID {
					fixtures = append(fixtures, fix)
				}
			}

			if venue == "home_away" {
				if fix.HomeTeam.ID == team.ID || fix.AwayTeam.ID == team.ID {
					fixtures = append(fixtures, fix)
				}
			}
		}
	}

	return fixtures
}
