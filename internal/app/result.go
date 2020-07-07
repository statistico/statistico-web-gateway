package app

type Result struct {
	ID       uint64      `json:"id"`
	HomeTeam Team        `json:"homeTeam"`
	AwayTeam Team        `json:"awayTeam"`
	Season   Season      `json:"season"`
	Round    Round       `json:"round"`
	Venue    Venue       `json:"venue"`
	DateTime JsonDate    `json:"date"`
	Stats    ResultStats `json:"stats"`
}

type ResultStats struct {
	HomeScore uint8 `json:"homeScore"`
	AwayScore uint8 `json:"awayScore"`
}
