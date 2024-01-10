package util

import (
	"testing"
	"time"

	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestRoundToFiveMinutes(t *testing.T) {
	rawTimestamp, _ := ConvertTimestampDefault("2024-03-14 20:17:00")
	rawTime := time.UnixMilli(rawTimestamp)
	roundedTime := RoundToFiveMinutes(rawTime)

	wantTime, _ := ConvertTimestampDefault("2024-03-14 20:15:00")
	assert.Assert(t, roundedTime.UnixMilli() == wantTime)

	rawTimestamp, _ = ConvertTimestampDefault("2024-03-14 20:18:00")
	rawTime = time.UnixMilli(rawTimestamp)
	roundedTime = RoundToFiveMinutes(rawTime)

	wantTime, _ = ConvertTimestampDefault("2024-03-14 20:20:00")
	assert.Assert(t, roundedTime.UnixMilli() == wantTime)
}

func TestRoundDownToFiveMinutes(t *testing.T) {
	rawTimestamp, _ := ConvertTimestampDefault("2024-03-14 20:17:00")
	rawTime := time.UnixMilli(rawTimestamp)
	roundedTime := RoundDownToFiveMinutes(rawTime)

	wantTime, _ := ConvertTimestampDefault("2024-03-14 20:15:00")
	assert.Assert(t, roundedTime.UnixMilli() == wantTime)

	rawTimestamp, _ = ConvertTimestampDefault("2024-03-14 20:18:00")
	rawTime = time.UnixMilli(rawTimestamp)
	roundedTime = RoundDownToFiveMinutes(rawTime)

	wantTime, _ = ConvertTimestampDefault("2024-03-14 20:15:00")
	assert.Assert(t, roundedTime.UnixMilli() == wantTime)
}

func TestRoundDownToFiveMinutesV2(t *testing.T) {
	rawTimestamp, _ := ConvertTimestampDefault("2024-03-14 20:17:00")
	roundedTime := RoundDownToFiveMinutesV2(rawTimestamp)

	wantTime, _ := ConvertTimestampDefault("2024-03-14 20:15:00")
	assert.Assert(t, roundedTime == wantTime)

	rawTimestamp, _ = ConvertTimestampDefault("2024-03-14 20:18:00")
	roundedTime = RoundDownToFiveMinutesV2(rawTimestamp)

	wantTime, _ = ConvertTimestampDefault("2024-03-14 20:15:00")
	assert.Assert(t, roundedTime == wantTime)
}
