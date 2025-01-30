package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Namith667/GoQuick/internal/config"
	"github.com/Namith667/GoQuick/internal/db"
	"github.com/Namith667/GoQuick/internal/logger"
	"github.com/Namith667/GoQuick/internal/routes"
	"go.uber.org/zap"
	//"github.com/joho/godotenv"
	//"gorm.io/driver/postgres"
)

func main() {

	logInstance := logger.Init()
	defer logInstance.Sync()

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
	logger.Log.Info("port is:: ", zap.String("port: ", port))
	if port == "" {
		port = "8080" // Default port
	}
	logger.Log.Info("Starting server", zap.String("address", port))
	err = http.ListenAndServe(port, r)
	if err != nil {
		logger.Log.Error("Server failed to start", zap.Error(err))
	}
}
