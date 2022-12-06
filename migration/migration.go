package migration

import (
	"log"

	"github.com/pressly/goose"
	"gorm.io/gorm"
)

func Up(db *gorm.DB) {

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Fail to connect db", err)
	}

	if err := goose.SetDialect("mysql"); err != nil {
		panic(err)
	}

	if err := goose.Run("up", sqlDB, "./migration"); err != nil {
		panic(err)
	}

}
