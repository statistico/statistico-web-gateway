package composer

import "time"

type ResultFilters struct {
	DateAfter  *time.Time  `json:"dateAfter"`
	DateBefore *time.Time  `json:"dateBefore"`
	Limit      *uint64     `json:"limit"`
	SeasonIds  *[]uint64   `json:"seasonIds"`
	Sort       *string     `json:"sort"`
	Team       *TeamFilter `json:"team"`
	Venue      *string     `json:"venue"`
}

type TeamFilter struct {
	ID    uint64  `json:"id"`
	Venue *string `json:"venue"`
}

type TeamStatFilters struct {
	DateAfter  *time.Time  `json:"dateAfter"`
	DateBefore *time.Time  `json:"dateBefore"`
	Limit      *uint64     `json:"limit"`
	Opponent   *bool       `json:"opponent"`
	SeasonIds  *[]uint64   `json:"seasonIds"`
	Sort       *string     `json:"sort"`
	Stat       string      `json:"stat"`
	Team       TeamFilter `json:"team"`
}
