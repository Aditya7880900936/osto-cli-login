package database

import (
	"fmt"

	"github.com/Aditya7880900936/osto-cli-login/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is the global database instance shared across the application.
var DB *gorm.DB

// ConnectDB establishes the SQLite database connection
// and performs automatic schema migration.
func ConnectDB() error {

	db, err := gorm.Open(sqlite.Open("storage/auth.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
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

// GetDB returns the initialized database instance.
func GetDB() *gorm.DB {
	return DB
}
