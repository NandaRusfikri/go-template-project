package pkg

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-template-project/docs"
	"os"
)

func InitSwagger(httpGin *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Title = "go-template-project"
	docs.SwaggerInfo.Description = "This Project Template Go which i usually use"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	//	@title			go-template-project
	//	@version		1.0
	//	@description	This Project Template Go which I usually use.
	//	@termsOfService	http://swagger.io/terms/

	//	@contact.name				API Support
	//	@contact.url				http://www.swagger.io/support
	//	@contact.email				nandarusfikri@gmail.com
	//	@license.name				Apache 2.0
	//	@license.url				http://www.apache.org/licenses/LICENSE-2.0.html
	//	@query.collection.format	multi
	//	@securityDefinitions.apikey	ApiKeyAuth
	//	@in							header
	//	@name						Authorization
	//	@x-extension-openapi		{"example": "value on a json format"}

	httpGin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
