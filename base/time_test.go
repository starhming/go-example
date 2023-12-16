package base

import (
	"fmt"
	"testing"
	"time"
)

const (
	SimpleTimeFormat           = "2006-01-02 15:04:05"
	SimpleTimeFormatWithRegion = "2006-01-02T15:04:05.000-07:00"
	TimeFormatDruid            = "2006-01-02T15:04:05+08:00"
	TimeFormatZ                = "2006-01-02T15:04:05.000Z"
	MyDateFormat               = "20060102"
	MyDateFormat1              = "2006-01-02"
	TimeThreshold              = 9999999999
)

func TestTimeFormat(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Format(MyDateFormat))
}
