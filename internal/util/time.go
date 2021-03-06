package util

import (
	"time"
)

func ProfileTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Debugf("%s [%s]", name, elapsed)
}
