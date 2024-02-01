package conversion

import (
	"time"

	"github.com/awlsring/texit/pkg/gen/texit"
)

func maybeMakeTime(t *time.Time) texit.OptFloat64 {
	if t == nil {
		return texit.OptFloat64{}
	}
	return texit.NewOptFloat64(float64(t.Unix()))
}

func maybeMakeFloat64(f float64) texit.OptFloat64 {
	if f == 0 {
		return texit.OptFloat64{}
	}
	return texit.NewOptFloat64(f)
}

func maybeMakeString(s string) texit.OptString {
	if s == "" {
		return texit.OptString{}
	}
	return texit.NewOptString(s)
}

func maybeMakeBool(b bool) texit.OptBool {
	if !b {
		return texit.OptBool{}
	}
	return texit.NewOptBool(b)
}
