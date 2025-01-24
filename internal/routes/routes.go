package routes

import (
	"log"

	"github.com/Namith667/GoQuick/internal/db"
	"github.com/Namith667/GoQuick/internal/handlers"
	"github.com/Namith667/GoQuick/internal/services"
	"github.com/gorilla/mux"
)

func InitRoutes(database db.Database) *mux.Router {
	r := mux.NewRouter()

	conn, err := database.Connect()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	productHandler := handlers.NewProductHandler(database)

	r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")

	//producr routes
	r.HandleFunc("/products", productHandler.GetAllProducts).Methods("GET")
	r.HandleFunc("/products/{id}", productHandler.GetProductById).Methods("GET")
	r.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")

	//auth service

	authService := services.NewAuthService(conn)
	authHandler := handlers.NewAuthHandler(authService)
	//auth routes
	r.HandleFunc("/register", authHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/login", authHandler.Login).Methods("POST")
	return r
}
