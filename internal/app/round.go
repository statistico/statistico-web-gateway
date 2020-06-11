package app

type Round struct {
	ID        uint64   `json:"id"`
	Name      string   `json:"name"`
	SeasonID  uint64   `json:"seasonId"`
	StartDate JsonDate `json:"startDate"`
	EndDate   JsonDate `json:"endDate"`
}
