package service

import (
	"alipaycloudrun-demo-for-go/util"
	"log"
)

type httpService struct {
}

func (h httpService) HttpHandler() string {
	log.Println("get index.html")
	version := util.GetEnvDefault("PUB_SERVICE_REVISION", "golang-demo")
	hostName := util.GetEnvDefault("HOSTNAME", "golang-demo")
	result := "欢迎使用云托管!&服务版本：" + version + "&实例主机：" + hostName
	return result
}

func NewHttpService() IHttpService {
	return &httpService{}
}
