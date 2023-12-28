package server

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go_trunk_based/constant"
	"go_trunk_based/database"
	"go_trunk_based/dto"
	authcontroller "go_trunk_based/module/auth/controller"
	authrepo "go_trunk_based/module/auth/repository"
	authusecase "go_trunk_based/module/auth/usecase"
	defaultcontroller "go_trunk_based/module/default/controller"
	productcontroller "go_trunk_based/module/product/controller"
	productrepo "go_trunk_based/module/product/repository"
	itemusecase "go_trunk_based/module/product/usecase"
	usercontroller "go_trunk_based/module/user/controller"
	userrepo "go_trunk_based/module/user/repository"
	userusecase "go_trunk_based/module/user/usecase"
	"go_trunk_based/pkg"
	"strconv"
)

func Start() {

	ConfEnv := pkg.LoadEnvironment(".env")

	if err := pkg.LoadFeature(".env.feature"); err != nil {
		panic(err)
	}
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

	if constant.FEATURE_MODULE_PRODUCT {
		productcontroller.InitProductControllerHTTP(httpServer, ItemUseCase)
	}

	defaultcontroller.InitDefaultController(httpServer)

	err = httpServer.Run(fmt.Sprintf(`:%v`, RESTPort))
	if err != nil {
		panic(err)
	}

}
