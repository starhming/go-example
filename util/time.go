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
