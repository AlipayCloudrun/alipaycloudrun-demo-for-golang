//go:build wireinject

package main

import (
	"alipaycloudrun-demo-for-go/api"
	"alipaycloudrun-demo-for-go/api/handlers"
	"alipaycloudrun-demo-for-go/app"
	"alipaycloudrun-demo-for-go/internal/service"
	"alipaycloudrun-demo-for-go/repo"
	"github.com/google/wire"
)

func InitServer() *app.Server {
	wire.Build(
		app.InitDB,
		app.InitRedis,
		api.NewRouter,
		handlers.NewHttpHandler,
		handlers.NewRecordHandler,
		handlers.NewRedisHandler,
		app.NewServer,
		app.NewGinEngine,
		service.NewHttpService,
		service.NewDbService,
		service.NewRedisService,
		repo.NewMysqlArticleRepository,
	)

	return &app.Server{}
}
