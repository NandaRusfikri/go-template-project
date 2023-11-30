package server

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-template-project/database"
	"go-template-project/dto"
	authcontroller "go-template-project/module/auth/controller"
	authrepo "go-template-project/module/auth/repository"
	authusecase "go-template-project/module/auth/usecase"
	defaultcontroller "go-template-project/module/default/controller"
	productcontroller "go-template-project/module/product/controller"
	productrepo "go-template-project/module/product/repository"
	itemusecase "go-template-project/module/product/usecase"
	usercontroller "go-template-project/module/user/controller"
	userrepo "go-template-project/module/user/repository"
	userusecase "go-template-project/module/user/usecase"
	"go-template-project/pkg"
	"strconv"
)

func Start() {

	ConfEnv := pkg.LoadEnvironment(".env")
	RESTPort, err := strconv.Atoi(ConfEnv.RestPort)
	if err != nil {
		log.Errorln("REST_PORT is not valid ", err.Error())
	}

	//redisClient := pkg.InitRedis(pkg.GetRedisConfig())
	//err = redisClient.New()
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}

	SMTPPort, err := strconv.Atoi(ConfEnv.SmtpPort)
	if err != nil {
		log.Errorln("SMTP Port is not valid ", err.Error())
	}

	smtpClient := pkg.InitEmail(&dto.SMTPConfig{
		Host:     ConfEnv.SmtpHost,
		Port:     SMTPPort,
		Email:    ConfEnv.SmtpEmail,
		Password: ConfEnv.SmtpPassword,
		Name:     ConfEnv.SmtpName,
	})

	DBPostgres, err := database.SetupDBPostgres(ConfEnv)
	if err != nil {
		log.Fatal(err.Error())
	}

	httpServer := pkg.SetupGin(ConfEnv)
	pkg.InitSwagger(httpServer)

	AuthRepo := authrepo.InitAuthRepository(DBPostgres)
	UserRepo := userrepo.InitUserRepository(DBPostgres)
	ProductRepo := productrepo.InitProductRepository(DBPostgres)
	ItemUseCase := itemusecase.InitProductUseCase(ProductRepo)
	AuthUseCase := authusecase.InitAuthUseCase(AuthRepo, UserRepo, smtpClient)
	UserUseCase := userusecase.InitUserUseCase(UserRepo)

	authcontroller.InitAuthControllerHTTP(httpServer, AuthUseCase)
	usercontroller.InitUserControllerHTTP(httpServer, UserUseCase)
	productcontroller.InitProductControllerHTTP(httpServer, ItemUseCase)
	defaultcontroller.InitDefaultController(httpServer)

	err = httpServer.Run(fmt.Sprintf(`:%v`, RESTPort))
	if err != nil {
		panic(err)
	}

}
