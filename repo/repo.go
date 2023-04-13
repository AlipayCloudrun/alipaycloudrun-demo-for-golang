package repo

import "alipaycloudrun-demo-for-go/models"

type IRecordRepo interface {
	GetRecordList() ([]models.RecordInfo, error)
	AddRecord(info *models.RecordInfo) error
	DeleteRecord(id int32) error
}
