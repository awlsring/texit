package conversion

import (
	"testing"
	"time"

	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/stretchr/testify/assert"
)

func TestMaybeMakeTime(t *testing.T) {
	now := time.Now()
	assert.Equal(t, texit.NewOptFloat64(float64(now.Unix())), maybeMakeTime(&now))

	var nilTime *time.Time
	assert.Equal(t, texit.OptFloat64{}, maybeMakeTime(nilTime))
}

func TestMaybeMakeFloat64(t *testing.T) {
	assert.Equal(t, texit.NewOptFloat64(123.45), maybeMakeFloat64(123.45))
	assert.Equal(t, texit.OptFloat64{}, maybeMakeFloat64(0))
}

func TestMaybeMakeString(t *testing.T) {
	assert.Equal(t, texit.NewOptString("test"), maybeMakeString("test"))
	assert.Equal(t, texit.OptString{}, maybeMakeString(""))
}

func TestMaybeMakeBool(t *testing.T) {
	assert.Equal(t, texit.NewOptBool(true), maybeMakeBool(true))
	assert.Equal(t, texit.OptBool{}, maybeMakeBool(false))
}
