package service

import (
	"alipaycloudrun-demo-for-go/models"
	"alipaycloudrun-demo-for-go/repo"
	"log"
)

type dbService struct {
	recordRepo repo.IRecordRepo
}

func NewDbService(recordRepo repo.IRecordRepo) IDbService {
	return &dbService{
		recordRepo: recordRepo,
	}
}

func (db *dbService) FetchList() (res []models.RecordInfo, err error) {
	log.Println("get record list info")
	return db.recordRepo.GetRecordList()
}

func (db *dbService) AddRecord(info *models.RecordInfo) error {
	log.Println("add record")
	return db.recordRepo.AddRecord(info)
}

func (db *dbService) DeleteRecord(id int32) error {
	log.Println("delete record")
	return db.recordRepo.DeleteRecord(id)
}
