package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	if db != nil {
		return db
	}

  var err error

	db, err = gorm.Open(sqlite.Open("cats.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}


	return db
}

func CloseDB() {
	if db == nil {
		return
	}

	dbInstance, err := db.DB()

	if err != nil {
		log.Println("Error getting db", err)
	}

	dbInstance.Close()

	db = nil
}
