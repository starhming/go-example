package util

import "time"

const (
	SimpleTimeFormat           = "2006-01-02 15:04:05"
	SimpleTimeFormatWithRegion = "2006-01-02T15:04:05.000-07:00"
	TimeFormatDruid            = "2006-01-02T15:04:05+08:00"
	TimeFormatZ                = "2006-01-02T15:04:05.000Z"
	MyDateFormat               = "20060102"
	MyDateFormat1              = "2006-01-02"
	TimeThreshold              = 9999999999
)

// ConvertTimestampDefault location Asia/Shanghai
func ConvertTimestampDefault(dateTime string) (int64, error) {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	location, err := time.ParseInLocation(SimpleTimeFormat, dateTime, cstSh)
	if err != nil {
		return 0, err
	}
	return location.UnixMilli(), err
}

// RoundToFiveMinutes 将时间舍入到最近的5分钟
func RoundToFiveMinutes(t time.Time) time.Time {
	return t.Round(5 * time.Minute)
}

// RoundDownToFiveMinutes 将时间向下舍入到最近的5分钟
func RoundDownToFiveMinutes(t time.Time) time.Time {
	// 计算最近的5分钟的整数值
	minutes := t.Minute() / 5 * 5

	// 构建新的时间对象
	roundedTime := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), minutes, 0, 0, t.Location())

	return roundedTime
}

// RoundDownToFiveMinutesV2 将时间向下舍入到最近的5分钟
func RoundDownToFiveMinutesV2(ts int64) int64 {
	fiveMinute := time.Minute * 5
	ts /= fiveMinute.Milliseconds()
	ts *= fiveMinute.Milliseconds()
	return ts
}
