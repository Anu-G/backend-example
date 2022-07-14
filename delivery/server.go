package delivery

import (
	"log"
	"os"

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

	UseCaseManager manager.UseCaseManagerInterface
}

func Server() *appServer {
	r := gin.Default()

	appCfg := config.NewConfig()
	dbCon := manager.NewInfraSetup(appCfg)
	repoManager := manager.NewRepo(dbCon)
	usecaseManager := manager.NewUseCase(repoManager)

	cfgServer := &appServer{
		engine:         r,
		host:           appCfg.APIConfig.APIUrl,
		UseCaseManager: usecaseManager,
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
	controller.NewMenuController(a.engine, a.UseCaseManager.MenuUseCase())
	controller.NewCustomerController(a.engine, a.UseCaseManager.CustomerUseCase())
	controller.NewTransactionController(a.engine, a.UseCaseManager.TrxUseCase())
	controller.NewTableController(a.engine, a.UseCaseManager.TableUseCase())
	controller.NewDiscountController(a.engine, a.UseCaseManager.DiscountUseCase())
}

func (a *appServer) Run() {
	if a.startServer {
		a.initControllers()
		if err := a.engine.Run(a.host); err != nil {
			panic(err)
		}
	}
}
