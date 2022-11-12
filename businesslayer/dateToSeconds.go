package businesslayer

import (
	"time"
)

func DateToSeconds(birthDate time.Time) int {
	today := time.Now()

	delta := today.Sub(birthDate)

	return int(delta.Seconds())
}
