package composer

import "time"

type Filters struct {
	Limit  *uint64   `json:"limit"`
	DateBefore   *time.Time   `json:"dateBefore"`
	DateAfter   *time.Time   `json:"dateAfter"`
	Venue     *string     `json:"venue"`
	SeasonIDs   []uint64   `json:"seasonIds"`
	Sort        *string    `json:"sort"`
}

