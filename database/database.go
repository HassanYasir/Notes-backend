package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Open SQLite connection
	db, err := gorm.Open(sqlite.Open("database/notes.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	DB = db
}
