package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Email        *string
	Age          uint8 `gorm:"default:18"`
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
}

// DiagnoseDataRecord 诊断过的流会把数据缓存至本地
// 查询sql：where stream_name = 'xx' and (did = 'x' or uid = 'x')
type DiagnoseDataRecord struct {
	gorm.Model
	StreamName   string
	RoomID       string
	UserIdentity string
	UserID       string
	DeviceID     string
	StartTime    time.Time
	EndTime      time.Time
	Status       int // -1: 数据只有离线数据，0: 在线数据，1: 数据已缓存至诊断es
}

func (ddr DiagnoseDataRecord) TableName() string {
	return "diagnose_data_records"
}
