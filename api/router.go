package api

import (
	"alipaycloudrun-demo-for-go/api/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	recordHandler handlers.RecordHandler
	httpHandler   handlers.HttpHandler
	redisHandler  handlers.RedisHandler
}

func NewRouter(recordHandler handlers.RecordHandler, httpHandler handlers.HttpHandler, redisHandler handlers.RedisHandler) *Router {
	router := &Router{
		recordHandler: recordHandler,
		httpHandler:   httpHandler,
		redisHandler:  redisHandler,
	}
	return router
}

func (r *Router) Index(engine *gin.Engine) {
	engine.Static("/static", "./static")
	engine.LoadHTMLFiles("./static/index.html")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
}

func (r *Router) Service(engine *gin.Engine) {
	engine.GET("/service", r.httpHandler.HttpTest)
}

func (r *Router) GetList(engine *gin.Engine) {
	engine.GET("/database/getList", r.recordHandler.GetList)
}

func (r *Router) AddRecord(engine *gin.Engine) {
	engine.POST("/database/addRecord", r.recordHandler.AddRecord)
}

func (r *Router) DeleteRecord(engine *gin.Engine) {
	engine.GET("/database/deleteRecord", r.recordHandler.DeleteRecord)
}

func (r *Router) RedisGet(engine *gin.Engine) {
	engine.GET("/redis/get", r.redisHandler.Get)
}

func (r *Router) RedisSet(engine *gin.Engine) {
	engine.GET("/redis/set", r.redisHandler.Set)
}
