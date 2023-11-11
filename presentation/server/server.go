package server

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-template-project/database"
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
	"strconv"
)

func Start() {

	ConfEnv := pkg.LoadEnvironment(".env")
	RESTPort, err := strconv.Atoi(ConfEnv.REST_PORT)
	if err != nil {
		log.Errorln("REST_PORT is not valid ", err.Error())
	}

	//redisClient := pkg.InitRedis(pkg.GetRedisConfig())
	//err = redisClient.New()
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}

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
	pkg.InitSwagger(httpGin)

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
