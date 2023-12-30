package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-template-project/constant"
	"go-template-project/dto"
	"go-template-project/util"
	"time"
)

type DefaultController struct {
	config dto.ConfigApp
}

func InitDefaultController(g *gin.Engine, config dto.ConfigApp) {
	handler := &DefaultController{
		config: config,
	}

	g.GET("/", handler.MainPage)
}

func (c *DefaultController) MainPage(ctx *gin.Context) {
	jsonData := map[string]interface{}{
		"service_name": c.config.ServiceName,
		"author":       constant.AUTHOR,
		"version":      constant.SERVICE_VERSION,
		"time_now":     time.Now(),
		"swagger":      fmt.Sprintf("http://%v/swagger/index.html", c.config.SwaggerHost),
	}
	util.APIResponse(ctx, dto.APIResponse{
		Code:    200,
		Message: "success",
		Data:    jsonData,
	})
}
