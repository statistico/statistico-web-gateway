package app

import (
	"fmt"
	"time"
)

type JsonDate time.Time

func (t JsonDate) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(time.RFC3339))
	return []byte(stamp), nil
}
