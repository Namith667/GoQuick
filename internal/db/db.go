package db

import (
	"log"

	"github.com/Namith667/GoQuick/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	dsn := "host=localhost user=testUser password=1234 dbname=GoQuickDB port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect")
	}
	return db
}

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Product{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate models", err)
	}

	log.Println("DB Migrations Sucess!!")
}
