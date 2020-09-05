package app

type TeamStat struct {
	FixtureID     uint64    `json:"fixtureId"`
	Result        *Result   `json:"result,omitempty"`
	Stat          string    `json:"stat"`
	Value         *uint32   `json:"value"`
}
