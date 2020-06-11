package composer

import (
	"github.com/statistico/statistico-web-gateway/internal/app"
	"time"
)

type FixtureSearch struct{}

func (f FixtureSearch) Search(q *FixtureSearchQuery) []*app.Fixture {
	fixtures := []*app.Fixture{}

	fix := &app.Fixture{
		ID: 1234,
		Competition: app.Competition{
			ID:    567,
			Name:  "English Premier League",
			IsCup: false,
		},
		Season: app.Season{
			ID:        89,
			Name:      "2019/2020",
			IsCurrent: true,
		},
		Round: app.Round{
			ID:        2,
			Name:      "2",
			SeasonID:  89,
			StartDate: app.JsonDate(time.Now()),
			EndDate:   app.JsonDate(time.Now()),
		},
		HomeTeam: app.Team{
			ID:   1,
			Name: "West Ham United",
		},
		AwayTeam: app.Team{
			ID:   10,
			Name: "Newcastle United",
		},
		Venue: app.Venue{
			ID:   4,
			Name: "London Stadium",
		},
		Date: app.JsonDate(time.Now()),
	}

	fixtures = append(fixtures, fix)

	return fixtures
}

type FixtureSearchQuery struct {
	LeagueIds         []uint64           `json:"leagueIds"`
	DateFrom          time.Time         `json:"dateTo"`
	DateTo            time.Time         `json:"dateFrom"`
	FixtureStatFilter FixtureStatFilter  `json:"statFilter"`
}

type FixtureStatFilter struct {
	Games   uint8     `json:"games"`
	Metric  string    `json:"metric"`
	Measure string    `json:"measure"`
	Team    string    `json:"team"`
	Type    string    `json:"type"`
	Value   float32   `json:"value"`
	Venue   string    `json:"venue"`
}
