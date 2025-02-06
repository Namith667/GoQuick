package routes

import (
	"github.com/Namith667/GoQuick/internal/db"
	"github.com/Namith667/GoQuick/internal/handlers"
	"github.com/Namith667/GoQuick/internal/logger"
	"github.com/Namith667/GoQuick/internal/middleware/auth"

	//"github.com/Namith667/GoQuick/internal/services"
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
	authService := auth.NewAuthService(conn)
	authHandler := handlers.NewAuthHandler(authService)

	//public routes
	r.Get("/health", handlers.HealthCheck)
	r.Get("/products", productHandler.GetAllProducts)
	r.Get("/products/{id}", productHandler.GetProductById)
	r.Post("/register", authHandler.RegisterUser)
	r.Post("/login", authHandler.Login)

	// Private (Authenticated) Routes
	r.Group(func(r chi.Router) {
		r.Use(auth.JWTAuthMiddleware)

		// Admin-only routes
		r.Group(func(r chi.Router) {
			r.Use(auth.RequireRole("admin"))
			r.Post("/products", productHandler.CreateProduct)
			r.Put("/products/{id}", productHandler.UpdateProduct)
			r.Delete("/products/{id}", productHandler.DeleteProduct)
		})
	})
	return r
}
