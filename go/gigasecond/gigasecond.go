package gigasecond

import (
	"math"
	"time"
)

const testVersion = 4

func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(math.Pow10(9)) * time.Second)
}
