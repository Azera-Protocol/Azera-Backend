package initializer

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error

	// github.com/mattn/go-sqlite3
	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB")
	}
}
