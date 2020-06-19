package app

type StatFilter struct {
	Action  string  `json:"action"`
	Games   uint8   `json:"games"`
	Measure string  `json:"measure"`
	Metric  string  `json:"metric"`
	Stat    string  `json:"stat"`
	Value   float32 `json:"value"`
	Venue   string  `json:"venue"`
}
