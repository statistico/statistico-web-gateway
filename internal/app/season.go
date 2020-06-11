package app

type Season struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	IsCurrent bool   `json:"isCurrent"`
}
