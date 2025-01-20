package routes

import (
	"github.com/Namith667/GoQuick/internal/handlers"
	"github.com/Namith667/GoQuick/internal/services"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func InitRoutes(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")

	//producr routes
	r.HandleFunc("/products", handlers.GetAllProducts).Methods("GET")
	r.HandleFunc("/products/{id}", handlers.GetProductById).Methods("GET")
	r.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")

	//auth service

	authService := services.NewAuthService(db)
	authHandler := handlers.NewAuthHandler(authService)
	//auth routes
	r.HandleFunc("/register", authHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/login", authHandler.Login).Methods("POST")
	return r
}
