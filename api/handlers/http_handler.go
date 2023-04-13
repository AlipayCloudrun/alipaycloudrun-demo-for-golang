package handlers

import (
	"alipaycloudrun-demo-for-go/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpHandler struct {
	HttpService service.IHttpService
}

func NewHttpHandler(httpService service.IHttpService) HttpHandler {
	return HttpHandler{
		HttpService: httpService,
	}
}

func (h *HttpHandler) HttpTest(c *gin.Context) {
	result := h.HttpService.HttpHandler()
	c.JSON(http.StatusOK, result)
}
