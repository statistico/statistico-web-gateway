package app

type Competition struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	IsCup bool   `json:"isCup"`
}
