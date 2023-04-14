package handlers

import (
	"alipaycloudrun-demo-for-go/internal/service"
	"alipaycloudrun-demo-for-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RecordHandler struct {
	dbService service.IDbService
}

func NewRecordHandler(dbService service.IDbService) RecordHandler {
	return RecordHandler{
		dbService: dbService,
	}
}

func (r *RecordHandler) GetList(c *gin.Context) {
	list, err := r.dbService.FetchList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    list,
	})
}

func (r *RecordHandler) AddRecord(c *gin.Context) {
	var recordInfo models.RecordInfo
	c.BindJSON(&recordInfo)

	err := r.dbService.AddRecord(&recordInfo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (r *RecordHandler) DeleteRecord(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"data":    err.Error(),
		})
		return
	}
	r.dbService.DeleteRecord(int32(id))
}
