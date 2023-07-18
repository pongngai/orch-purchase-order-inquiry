package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func New() *gorm.DB {
	var gormLogger logger.Interface

	gormLogger = logger.Default

	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=localhost user=postgres password=34647de759 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Bangkok")), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
