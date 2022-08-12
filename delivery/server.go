package delivery

import (
	"log"
	"os"

	"wmb-rest-api/auth"
	"wmb-rest-api/config"
	"wmb-rest-api/delivery/controller"
	"wmb-rest-api/manager"
	"wmb-rest-api/tool"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	engine      *gin.Engine
	host        string
	startServer bool

	UseCaseManager    manager.UseCaseManagerInterface
	MiddlewareManager manager.MiddlewareManagerInterface
	Auth              auth.TokenInterface
}

func Server() *appServer {
	r := gin.Default()
	r.Use(CORSMiddleware())

	appCfg := config.NewConfig()
	dbCon := manager.NewInfraSetup(appCfg)
	auth := auth.NewTokenService(appCfg.TokenConfig)
	repoManager := manager.NewRepo(dbCon)
	usecaseManager := manager.NewUseCase(repoManager)
	middlewareManager := manager.NewMiddleware(auth)

	cfgServer := &appServer{
		engine:            r,
		host:              appCfg.APIConfig.APIUrl,
		UseCaseManager:    usecaseManager,
		MiddlewareManager: middlewareManager,
		Auth:              auth,
	}

	args := os.Args[1:]
	if len(args) > 0 {
		switch args[0] {
		case "db:migrate":
			if appCfg.DBConfig.Environment == "DEV" {
				tool.RunMigrate(dbCon)
			} else {
				log.Fatal("not a dev env, cannot migrate")
			}
			return cfgServer
		case "db:seeds":
			if appCfg.DBConfig.Environment == "DEV" {
				tool.RunSeeds(dbCon)
			} else {
				log.Fatal("not a dev env, cannot migrate")
			}
			return cfgServer
		default:
			log.Fatalln("argument not found !")
			cfgServer.startServer = true
			return cfgServer
		}
	}

	cfgServer.startServer = true
	return cfgServer
}

func (a *appServer) initControllers() {
	controller.NewMenuController(a.engine, a.UseCaseManager.MenuUseCase(), a.MiddlewareManager.AuthMiddleware())
	controller.NewCustomerController(a.engine, a.UseCaseManager.CustomerUseCase(), a.MiddlewareManager.AuthMiddleware())
	controller.NewTransactionController(a.engine, a.UseCaseManager.TrxUseCase(), a.MiddlewareManager.AuthMiddleware())
	controller.NewTableController(a.engine, a.UseCaseManager.TableUseCase(), a.MiddlewareManager.AuthMiddleware())
	controller.NewDiscountController(a.engine, a.UseCaseManager.DiscountUseCase(), a.MiddlewareManager.AuthMiddleware())
	controller.NewAuthController(a.engine, a.UseCaseManager.CustomerUseCase(), a.UseCaseManager.AuthUseCase(), a.Auth)
}

func (a *appServer) Run() {
	if a.startServer {
		a.initControllers()
		if err := a.engine.Run(a.host); err != nil {
			panic(err)
		}
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(206)
			return
		}

		c.Next()
	}
}
