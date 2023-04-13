package handlers

import (
	"alipaycloudrun-demo-for-go/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RedisHandler struct {
	redisService service.IRedisService
}

func NewRedisHandler(redisService service.IRedisService) RedisHandler {
	return RedisHandler{
		redisService: redisService,
	}
}

func (r *RedisHandler) Set(ctx *gin.Context) {
	key := ctx.Query("key")
	value := ctx.Query("value")

	err := r.redisService.Set(key, value)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    "redis set fail, reason is " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (r *RedisHandler) Get(ctx *gin.Context) {
	key := ctx.Query("key")

	value, err := r.redisService.Get(key)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    "redis get fail, reason is " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    value,
	})
}
