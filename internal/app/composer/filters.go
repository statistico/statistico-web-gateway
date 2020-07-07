package composer

import "time"

type Filters struct {
	Limit  *uint64   `json:"limit"`
	DateBefore   *time.Time   `json:"dateBefore"`
	DateAfter   *time.Time   `json:"dateAfter"`
	SeasonIDs   []uint64   `json:"seasonIds"`
	Sort        *string    `json:"sort"`
	TeamIDs   []uint64     `json:"teamIds"`
	Venue     *string     `json:"venue"`
}

