package composer

import "time"

type Filters struct {
	Limit  *uint64   `json:"limit"`
	DateBefore   *time.Time   `json:"dateBefore"`
	DateAfter   *time.Time   `json:"dateAfter"`
	Sort        *string    `json:"sort"`
	Venue     *string     `json:"venue"`
}

