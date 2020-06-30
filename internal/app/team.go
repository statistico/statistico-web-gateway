package app

type Team struct {
	ID           uint64  `json:"id"`
	Name         string  `json:"name"`
	ShortCode    *string `json:"shortCode"`
	CountryID    uint64  `json:"countryId"`
	VenueID      uint64  `json:"venueId"`
	NationalTeam bool    `json:"nationalTeam"`
	Founded      *uint64 `json:"founded"`
	Logo         *string `json:"logo"`
}
