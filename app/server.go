package app

import (
	"alipaycloudrun-demo-for-go/api"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine    *gin.Engine
	apiRouter *api.Router
}

func NewServer(engin *gin.Engine, apiRouter *api.Router) *Server {
	return &Server{
		engine:    engin,
		apiRouter: apiRouter,
	}
}

func (s *Server) Start() {
	s.apiRouter.Index(s.engine)
	s.apiRouter.Service(s.engine)
	s.apiRouter.GetList(s.engine)
	s.apiRouter.AddRecord(s.engine)
	s.apiRouter.DeleteRecord(s.engine)
	s.apiRouter.RedisGet(s.engine)
	s.apiRouter.RedisSet(s.engine)
	s.engine.Run("")
}
