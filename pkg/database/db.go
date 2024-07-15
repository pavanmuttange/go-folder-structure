package database

import (
	"log"

	"github.com/pavanmuttange/go-folder-structure/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	db.AutoMigrate(&models.User{})

	return db
}
