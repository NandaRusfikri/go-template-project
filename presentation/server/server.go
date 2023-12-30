package server

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-template-project/database"
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

	config := pkg.LoadConfig(".env")
	restPort, err := strconv.Atoi(config.App.RestPort)
	if err != nil {
		log.Errorln("REST_PORT is not valid ", err.Error())
	}

	//redisClient := pkg.InitRedis(pkg.GetRedisConfig())
	//err = redisClient.New()
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}

	smtpClient := pkg.InitEmail(&config.SMTP)

	dbPostgres, err := database.InitDBPostgres(config.Database, config.App.Timezone)
	if err != nil {
		log.Fatal(err.Error())
	}

	httpServer := pkg.InitHTTPGin(config.App)
	pkg.InitSwagger(httpServer)

	authRepo := authrepo.InitAuthRepository(dbPostgres)
	userRepo := userrepo.InitUserRepository(dbPostgres)
	productRepo := productrepo.InitProductRepository(dbPostgres)
	productUseCase := itemusecase.InitProductUseCase(productRepo)
	authUseCase := authusecase.InitAuthUseCase(authRepo, userRepo, smtpClient)
	userUseCase := userusecase.InitUserUseCase(userRepo)

	authcontroller.InitAuthControllerHTTP(httpServer, authUseCase)
	usercontroller.InitUserControllerHTTP(httpServer, userUseCase)
	productcontroller.InitProductControllerHTTP(httpServer, productUseCase)
	defaultcontroller.InitDefaultController(httpServer)

	err = httpServer.Run(fmt.Sprintf(`:%v`, restPort))
	if err != nil {
		panic(err)
	}

}
