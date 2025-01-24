package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Namith667/GoQuick/internal/config"
	"github.com/Namith667/GoQuick/internal/db"
	"github.com/Namith667/GoQuick/internal/routes"
	//"github.com/joho/godotenv"
	//"gorm.io/driver/postgres"
)

func main() {

	config.LoadEnv()

	dsn := config.GetDSN()
	database := db.NewPostgresDB(dsn)

	conn, err := database.Connect()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return
	}

	fmt.Println("Database connected successfully!", conn)
	// Run migrations
	database.RunMigrations()

	r := routes.InitRoutes(database)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}
	fmt.Println("starting server on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

}

// /internal
// 	/config
// 		config.go
// 	/db
// 		db.go
// 		postgres.go
// 	/handlers
// 	/models
// 	/routes
// 	/services
// .env
// main.go
