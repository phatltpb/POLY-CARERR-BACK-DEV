package db

import (
	"fmt"
	"log"
	"time"

	"github.com/tuongnguyen1209/poly-career-back/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDb() {
	var (
		config = config.GetConfig()
		host   = config.Mysql.Host
		port   = config.Mysql.Port
		user   = config.Mysql.User
		pass   = config.Mysql.Pass
		dbName = config.Mysql.Database
	)

	connectionString := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		pass,
		host,
		port,
		dbName,
	)

	fmt.Println(connectionString)

	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
	}
	dbConnect, err := gorm.Open(mysql.Open(connectionString), gormConfig)

	if err != nil {
		log.Fatal("Fail to connect db", err)
	}

	log.Println("Connected db", dbConnect)

	db = dbConnect.Debug()
	rawDB, _ := db.DB()
	rawDB.SetMaxIdleConns(10)
	rawDB.SetMaxOpenConns(44)
	rawDB.SetConnMaxLifetime(time.Minute * 5)
}

func GetDb() *gorm.DB {
	return db
}

func CloseConnect(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	fmt.Println("close connection")
	dbSQL.Close()
}
