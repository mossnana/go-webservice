package db

import (
	"todo/webservice/services/config"
	"todo/webservice/services/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB(envConfig *config.ENVConfig, log *log.STDLogger) (db *gorm.DB) {
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	db, e := gorm.Open(postgres.Open(envConfig.DB), gormConfig)
	if e != nil {
		panic(e)
	}

	log.Info("connect database success")

	return
}
