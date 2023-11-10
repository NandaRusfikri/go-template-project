package server

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-template-project/database"
	"go-template-project/docs"
	auth_controller "go-template-project/module/auth/controller"
	auth_repo "go-template-project/module/auth/repository"
	auth_usecase "go-template-project/module/auth/usecase"
	default_controller "go-template-project/module/default/controller"
	product_controller "go-template-project/module/product/controller"
	product_repo "go-template-project/module/product/repository"
	item_usecase "go-template-project/module/product/usecase"
	user_controller "go-template-project/module/user/controller"
	user_repo "go-template-project/module/user/repository"
	user_usecase "go-template-project/module/user/usecase"
	"go-template-project/pkg"
	"os"
	"strconv"
)

func Start() {

	ConfEnv := pkg.LoadEnvironment(".env")
	RESTPort, err := strconv.Atoi(ConfEnv.REST_PORT)
	if err != nil {
		log.Errorln("REST_PORT is not valid ", err.Error())
	}

	SMTPPort, err := strconv.Atoi(ConfEnv.SMTP_PORT)
	if err != nil {
		log.Errorln("SMTP Port is not valid ", err.Error())
	}

	smtpClient := pkg.InitEmail(&pkg.SMTPConfig{
		Host:     ConfEnv.SMTP_HOST,
		Port:     SMTPPort,
		Email:    ConfEnv.SMTP_EMAIL,
		Password: ConfEnv.SMTP_PASSWORD,
		Name:     ConfEnv.SMTP_NAME,
	})

	DBPostgres, err := database.SetupDBPostgres(ConfEnv)
	if err != nil {
		log.Fatal(err.Error())
	}

	httpGin := pkg.SetupGin(ConfEnv)
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Title = "go-template-project"
	docs.SwaggerInfo.Description = "go-template-project"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// @title go-template-project
	// @version 1.0
	// @description This is a sample server celler server.
	// @termsOfService http://swagger.io/terms/

	// @contact.name API Support
	// @contact.url http://www.swagger.io/support
	// @contact.email nandarusfikri@gmail.com
	// @license.name Apache 2.0
	// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
	// @query.collection.format multi
	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization
	// @x-extension-openapi {"example": "value on a json format"}

	httpGin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	AuthRepo := auth_repo.InitAuthRepository(DBPostgres)
	UserRepo := user_repo.InitUserRepository(DBPostgres)
	ProductRepo := product_repo.InitProductRepository(DBPostgres)
	ItemUseCase := item_usecase.InitProductUseCase(ProductRepo)
	AuthUseCase := auth_usecase.InitAuthUseCase(AuthRepo, UserRepo, smtpClient)
	UserUseCase := user_usecase.InitUserUseCase(UserRepo)

	auth_controller.InitAuthControllerHTTP(httpGin, AuthUseCase)
	user_controller.InitUserControllerHTTP(httpGin, UserUseCase)
	product_controller.InitProductControllerHTTP(httpGin, ItemUseCase)
	default_controller.InitDefaultController(httpGin)

	httpGin.Run(fmt.Sprintf(`:%v`, RESTPort))

}
