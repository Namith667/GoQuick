package routes

import (
	"github.com/Namith667/GoQuick/internal/db"
	"github.com/Namith667/GoQuick/internal/handlers"
	"github.com/Namith667/GoQuick/internal/logger"
	"github.com/Namith667/GoQuick/internal/middleware"
	"github.com/Namith667/GoQuick/internal/services"
	"go.uber.org/zap"

	//"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5"
)

func InitRoutes(database db.Database) *chi.Mux {
	r := chi.NewRouter()

	conn, err := database.Connect()
	if err != nil {
		logger.Log.Error("Database connection failed:", zap.Error(err))
	}

	productHandler := handlers.NewProductHandler(database)

	r.Get("/health", handlers.HealthCheck)

	//public routes
	r.Get("/products", productHandler.GetAllProducts)
	r.Get("/products/{id}", productHandler.GetProductById)

	//private routes
	// Product routes
	r.Get("/products", productHandler.GetAllProducts)
	r.Get("/products/{id}", productHandler.GetProductById)
	r.With(middleware.RequireRole("admin")).Post("/products", productHandler.CreateProduct)
	r.With(middleware.RequireRole("admin")).Put("/products/{id}", productHandler.UpdateProduct)
	r.With(middleware.RequireRole("admin")).Delete("/products/{id}", productHandler.DeleteProduct)

	// Auth routes
	authService := services.NewAuthService(conn)
	authHandler := handlers.NewAuthHandler(authService)
	r.Post("/register", authHandler.RegisterUser)
	r.Post("/login", authHandler.Login)

	return r
}
