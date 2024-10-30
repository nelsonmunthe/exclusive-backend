package db

import (
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GormPostgres(dsn string) *gorm.DB {
	logLevel, _ := strconv.Atoi(os.Getenv("DB_LOG_MODE"))
	if logLevel <= 0 {
		logLevel = 2
	}
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(logLevel)),
	})
	return db
}
