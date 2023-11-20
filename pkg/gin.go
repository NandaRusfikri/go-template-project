package pkg

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"go-template-project/dto"
)

func SetupGin(config dto.ConfigEnvironment) *gin.Engine {

	app := gin.Default()

	if config.GO_ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else if config.GO_ENV == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)

	}

	app.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(gzip.Gzip(gzip.DefaultCompression))

	return app
}
