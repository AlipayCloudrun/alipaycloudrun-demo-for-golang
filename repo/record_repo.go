package repo

import (
	"alipaycloudrun-demo-for-go/models"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type mysqlRecordRepository struct {
	DB *gorm.DB
}

func (m mysqlRecordRepository) GetRecordList() ([]models.RecordInfo, error) {
	var records []models.RecordInfo
	var err error
	if m.DB == nil {
		return nil, fmt.Errorf("数据库连接为空，请检查环境变量")
	}
	err = m.DB.Find(&records).Error
	return records, err
}

func (m mysqlRecordRepository) AddRecord(info *models.RecordInfo) error {
	if m.DB == nil {
		return fmt.Errorf("数据库连接为空，请检查环境变量")
	}
	info.GmtModify = time.Now()
	info.GmtCreate = time.Now()
	return m.DB.Create(info).Error
}

func (m mysqlRecordRepository) DeleteRecord(id int32) error {
	if m.DB == nil {
		return fmt.Errorf("数据库连接为空，请检查环境变量")
	}
	return m.DB.Delete(&models.RecordInfo{}, id).Error
}

func NewMysqlArticleRepository(DB *gorm.DB) IRecordRepo {
	return &mysqlRecordRepository{DB}
}
