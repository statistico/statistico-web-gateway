package app

type Fixture struct {
	ID          uint64      `json:"id"`
	Competition Competition `json:"competition"`
	Season      Season      `json:"season"`
	Round       Round       `json:"round"`
	HomeTeam    Team        `json:"homeTeam"`
	AwayTeam    Team        `json:"awayTeam"`
	Venue       Venue       `json:"venue"`
	Date        JsonDate    `json:"date"`
}
