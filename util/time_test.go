package util

import (
	"fmt"
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

func TestRFCTime(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Format(time.DateTime))
	fmt.Println(now.Format(time.RFC3339))

	fmt.Println("---------------------------------------------")

	// 定义时间字符串
	timeStr := "2024-08-06T16:43:04+08:00"

	// 解析时间
	ts, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		fmt.Println("解析时间出错：", err)
		return
	}

	// 打印解析后的时间
	fmt.Println("解析后的时间：", ts)

	fmt.Println("-------------------------------------------------")

	timeStr = "2024-08-19T13:03:00Z"
	ts, err = time.ParseInLocation(time.RFC3339, timeStr, time.Local)
	fmt.Println(ts.Format(time.DateTime))

}
