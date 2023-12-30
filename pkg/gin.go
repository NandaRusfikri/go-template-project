package pkg

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"go-template-project/dto"
)

func InitHTTPGin(config dto.ConfigApp) *gin.Engine {

	app := gin.Default()

	if config.GoENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else if config.GoENV == "test" {
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

type Middleware struct {
	MiddlewareConfig *cors.Config
}

type AppConfig struct {
	Debug        bool
	Host         string
	Port         string
	UseTLS       bool
	CertFilePath string
	KeyFilePath  string
	Cors         *cors.Config
}

func InitMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) CORS(appCfg *AppConfig) cors.Config {
	var httpCors cors.Config
	if appCfg.Debug {
		httpCors = cors.DefaultConfig()
		httpCors.AllowAllOrigins = true
		httpCors.AllowCredentials = true
		httpCors.AddAllowHeaders("authorization")
	} else {
		httpCors = cors.Config{
			AllowAllOrigins:        appCfg.Cors.AllowAllOrigins,
			AllowOrigins:           appCfg.Cors.AllowOrigins,
			AllowMethods:           appCfg.Cors.AllowMethods,
			AllowHeaders:           appCfg.Cors.AllowHeaders,
			ExposeHeaders:          appCfg.Cors.ExposeHeaders,
			AllowCredentials:       appCfg.Cors.AllowCredentials,
			AllowWildcard:          appCfg.Cors.AllowWildcard,
			AllowBrowserExtensions: appCfg.Cors.AllowBrowserExtensions,
			AllowWebSockets:        appCfg.Cors.AllowWebSockets,
			AllowFiles:             appCfg.Cors.AllowFiles,
			MaxAge:                 appCfg.Cors.MaxAge,
		}
	}

	return httpCors
}
