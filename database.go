package main

import "github.com/jinzhu/gorm"

func initDatabase() *gorm.DB {
	db, err := gorm.Open("sqlite3", "database/test.db")

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&Transaction{})
	db.AutoMigrate(&Event{})

	return db
}
