package app

type TeamStat struct {
	FixtureID     uint64    `json:"fixture_id"`
	Stat          string    `json:"stat"`
	Value         *uint32   `json:"value"`
}
