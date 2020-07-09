package composer

import "time"

type Filters struct {
	Limit      *uint64    `json:"limit"`
	DateBefore *time.Time `json:"dateBefore"`
	DateAfter  *time.Time `json:"dateAfter"`
	SeasonIds  *[]uint64  `json:"seasonIds"`
	Sort       *string    `json:"sort"`
	Team       *TeamFilter `json:"team"`
	Venue      *string    `json:"venue"`
}

type TeamFilter struct {
	ID   uint64    `json:"id"`
	Venue   *string  `json:"venue"`
}
