package database

import (
	"fmt"

	"github.com/Aditya7880900936/osto-cli-login/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {

	db, err := gorm.Open(sqlite.Open("storage/auth.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	DB = db

	return nil
}

func GetDB() *gorm.DB {
	return DB
}