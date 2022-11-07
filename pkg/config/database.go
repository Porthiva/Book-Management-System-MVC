package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		ENV.DB_USERNAME,
		ENV.DB_PASSWORD,
		ENV.DB_HOST,
		ENV.DB_PORT,
		ENV.DB_NAME,
	)
	dbLog := logger.Default.LogMode(logger.Info)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: dbLog})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
