package repo

import (
	"alipaycloudrun-demo-for-go/models"
	"gorm.io/gorm"
	"time"
)

type mysqlRecordRepository struct {
	DB *gorm.DB
}

func (m mysqlRecordRepository) GetRecordList() ([]models.RecordInfo, error) {
	var records []models.RecordInfo
	var err error
	err = m.DB.Find(&records).Error
	return records, err
}

func (m mysqlRecordRepository) AddRecord(info *models.RecordInfo) error {
	info.GmtModify = time.Now()
	info.GmtCreate = time.Now()
	return m.DB.Create(info).Error
}

func (m mysqlRecordRepository) DeleteRecord(id int32) error {
	return m.DB.Delete(&models.RecordInfo{}, id).Error
}

func NewMysqlArticleRepository(DB *gorm.DB) IRecordRepo {
	return &mysqlRecordRepository{DB}
}
