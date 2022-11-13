package businesslayer

import (
	"fmt"
	"time"
)

func DateToSeconds(birthDate time.Time) (int, error) {
	var err error = nil
	today := time.Now()

	delta := today.Sub(birthDate)
	if delta < 0 {
		err = fmt.Errorf("birthdate must be in the past")
	}

	return int(delta.Seconds()), err
}
