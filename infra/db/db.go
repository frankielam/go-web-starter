package db

import (
	"gowebstarter/infra/config"
	"log"
	"sync"

	// "gorm.io/driver/mysql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

func GetDB() *gorm.DB {
	dbOnce.Do(postgresInit)
	// dbOnce.Do(mysqlInit)	// if using mysql database
	return db
}

// Postgres Database
func postgresInit() {
	var err error
	dsn := config.GetConf().Db.Dsn
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Postgres connected")
}

// MySQL database
// func mysqlInit() {
// 	var err error
// 	dsn := config.GetConf().Db.Dsn
// 	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	log.Println("Mysql connected")
// }
