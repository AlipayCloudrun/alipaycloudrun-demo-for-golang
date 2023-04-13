package models

import "time"

type RecordInfo struct {
	Id        int32     `gorm:"column:id" json:"id"`
	Record    string    `gorm:"column:record" json:"record"`
	GmtCreate time.Time `gorm:"column:gmt_create" json:"gmtCreate"`
	GmtModify time.Time `gorm:"column:gmt_modified" json:"gmtModify"`
}
