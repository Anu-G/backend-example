package manager

import (
	"fmt"
	"log"
	"os"
	"time"

	"wmb-rest-api/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConnectionInterface interface {
	DBCon() *gorm.DB
}

type database struct {
	db  *gorm.DB
	cfg config.DBConfig
}

func NewDBConnection(config config.DBConfig) DBConnectionInterface {
	dbs := new(database)
	dbs.cfg = config
	dbs.db = dbs.dbConnect()
	return dbs
}

func (db *database) dbConnect() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v timezone=%v",
		db.cfg.DBHost, db.cfg.DBUser, db.cfg.DBPassword, db.cfg.DBName, db.cfg.DBPort, db.cfg.SSLMode, db.cfg.TimeZone)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
			Colorful: true,
		},
	)

	gormCfg := new(gorm.Config)
	if db.cfg.Environment == "DEV" {
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

func (db *database) DBCon() *gorm.DB {
	return db.db
}
