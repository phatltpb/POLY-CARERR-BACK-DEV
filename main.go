package main

import (
	"log"

	"github.com/tuongnguyen1209/poly-career-back/apis/delivery"
	"github.com/tuongnguyen1209/poly-career-back/config"
	"github.com/tuongnguyen1209/poly-career-back/db"
	"github.com/tuongnguyen1209/poly-career-back/migration"
)

func main() {
	db.ConnectDb()

	database := db.GetDb()
	defer func() {
		db.CloseConnect(database)
	}()

	migration.Up(database)
	e := delivery.Start(database)
	port := config.GetConfig().Port
	log.Fatal(e.Start(":" + port))
}
