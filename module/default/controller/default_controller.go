package controller

import (
	"github.com/gin-gonic/gin"
	"go-template-project/constant"
	"go-template-project/util"
	"net/http"
	"time"
)

type DefaultController struct{}

func InitDefaultController(g *gin.Engine) {
	handler := &DefaultController{}

	g.GET("/", handler.MainPage)
}

func (h *DefaultController) MainPage(c *gin.Context) {
	jsonData := map[string]interface{}{
		"application_name": constant.SERVICE_NAME,
		"author":           constant.AUTHOR,
		"version":          constant.SERVICE_VERSION,
		"time_server":      time.Now(),
		"swagger":          "http://localhost:54321/swagger/index.html",
	}
	util.APIResponse(c, "success", http.StatusOK, 0, jsonData)
}
