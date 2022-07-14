package manager

import (
	"fmt"
	"log"
	"os"
	"time"

	"wmb-rest-api/config"
	"wmb-rest-api/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type InfraManagerInterface interface {
	DBCon() *gorm.DB
	LopeiCon() service.LopeiPaymentClient
}

type infraManager struct {
	db          *gorm.DB
	lopeiClient service.LopeiPaymentClient

	cfgDB    config.DBConfig
	cfgLopei config.LopeiGrpcConfig
}

func NewInfraSetup(config config.Config) InfraManagerInterface {
	newInfra := new(infraManager)
	newInfra.cfgDB = config.DBConfig
	newInfra.cfgLopei = config.LopeiGrpcConfig
	newInfra.db = newInfra.dbConnect()
	newInfra.initLopei()
	return newInfra
}

func (im *infraManager) dbConnect() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v timezone=%v",
		im.cfgDB.DBHost, im.cfgDB.DBUser, im.cfgDB.DBPassword, im.cfgDB.DBName, im.cfgDB.DBPort, im.cfgDB.SSLMode, im.cfgDB.TimeZone)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
			Colorful: true,
		},
	)

	gormCfg := new(gorm.Config)
	if im.cfgDB.Environment == "DEV" {
		gormCfg.Logger = newLogger
	}

	dbcon, err := gorm.Open(postgres.Open(dsn), gormCfg)
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, _ := dbcon.DB()
	defer sqlDB.SetConnMaxLifetime(15 * time.Minute)

	return dbcon
}

func (im *infraManager) initLopei() {
	conn, err := grpc.Dial(im.cfgLopei.LopeiUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect lopei: %s", err)
	}
	client := service.NewLopeiPaymentClient(conn)
	im.lopeiClient = client
}

func (im *infraManager) DBCon() *gorm.DB {
	return im.db
}

func (im *infraManager) LopeiCon() service.LopeiPaymentClient {
	return im.lopeiClient
}
