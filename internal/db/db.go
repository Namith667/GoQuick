package db

import (
	"fmt"
	"log"
	"os"

	"github.com/Namith667/GoQuick/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	//dsn := "host=localhost user=testuser password=1234 dbname=goquickdb port=5432 sslmode=disable"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("SB_SSL_MODE"),
	)

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
