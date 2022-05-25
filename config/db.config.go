package config

import (
	"log"
	"os"

	"golang-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Set Database Connection
func SetupDatabaseConnection() *gorm.DB {
	dbURI := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(
		&models.Book{},
		&models.User{},
	)

	return db
}

// Close Database Connection
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()

	if err != nil {
		log.Fatalln(err)
	}

	dbSQL.Close()
}
