package base

import (
	"fmt"
	"testing"
	"time"

	"github.com/starshm/go-example/util"
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

func TestMindTime(t *testing.T) {
	now := time.Now()
	start := now
	end := now.Add(time.Hour * 5)

	mid := (start.UnixMilli() + end.UnixMilli()) >> 1

	formatPrintTime(start)
	formatPrintTime(time.UnixMilli(mid))
	formatPrintTime(end)

}

func formatPrintTime(t time.Time) {
	fmt.Println(t.Format(SimpleTimeFormat))
}

func TestParseTime(t *testing.T) {
	t1, _ := time.Parse(SimpleTimeFormat, "2023-11-16 19:06:00")
	fmt.Println(t1.UnixMilli())

	t2, _ := time.Parse(time.RFC3339, "2024-03-13T11:27:00+08:00")
	fmt.Println(t2.UnixMilli())

	timestampDefault, _ := util.ConvertTimestampDefault("2023-11-16 19:06:00")
	fmt.Println(timestampDefault)
}

func TestFormatTime(t *testing.T) {
	tenHourMS := time.Hour.Milliseconds() * 10
	fmt.Println(tenHourMS)
}
