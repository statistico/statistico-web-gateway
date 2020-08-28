package app

type TeamStat struct {
	FixtureID     uint64    `json:"fixtureId"`
	Stat          string    `json:"stat"`
	Value         *uint32   `json:"value"`
}
