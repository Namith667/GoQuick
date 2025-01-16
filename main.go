package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Namith667/GoQuick/internal/db"
	"github.com/Namith667/GoQuick/internal/routes"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error loading it")
	}

	database := db.Connect()

	db.RunMigrations(database)
	r := routes.InitRoutes()

	fmt.Println("starting Server :8080")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
