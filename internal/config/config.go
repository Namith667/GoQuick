package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Cannot load .env file", err)
	}
}

func GetDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL_MODE"),
	)
}

func GetExpirationTime() int {
	expTimeStr := os.Getenv("JWT_EXPIRATION_TIME")
	if expTimeStr == "" {
		log.Println("Warning: JWT_EXPIRATION_TIME not set, using default (24 hours).")
		return 24 // Default to 24 hours if not set
	}

	expTime, err := strconv.Atoi(expTimeStr)
	if err != nil {
		log.Println("Error parsing JWT_EXPIRATION_TIME, using default (24 hours).", err)
		return 24
	}

	return expTime
}
