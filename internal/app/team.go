package app

type Team struct {
	ID           uint64  `json:"id"`
	Name         string  `json:"name"`
	ShortCode    *string `json:"short_code"`
	CountryID    uint64  `json:"country_id"`
	VenueID      uint64  `json:"venue_id"`
	NationalTeam bool    `json:"national_team"`
	Founded      *uint64 `json:"founded"`
	Logo         *string `json:"logo"`
}
