package service

import "alipaycloudrun-demo-for-go/models"

type IHttpService interface {
	HttpHandler() string
}

type IDbService interface {
	FetchList() (res []models.RecordInfo, err error)
	AddRecord(info *models.RecordInfo) error
	DeleteRecord(id int32) error
}

type IRedisService interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
}
