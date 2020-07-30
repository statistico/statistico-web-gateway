package app

type Competition struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	IsCup     bool   `json:"isCup"`
	CountryID uint64 `json:"countryId"`
}
