package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Namith667/GoQuick/internal/db"
	"github.com/Namith667/GoQuick/internal/logger"
	"github.com/Namith667/GoQuick/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// ProductHandler struct with dependency injection
type ProductHandler struct {
	DB db.Database
}

// NewProductHandler initializes ProductHandler
func NewProductHandler(database db.Database) *ProductHandler {
	return &ProductHandler{DB: database}
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {

	var products []models.Product
	conn, err := h.DB.Connect()
	if err != nil {
		logger.Log.Error("DB Connection error", zap.Error(err))
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	conn.Find(&products)
	logger.Log.Info("Product fetch success")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(products)

}

func (h *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	conn, err := h.DB.Connect()
	if err != nil {
		logger.Log.Error("DB Connection error", zap.Error(err))
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}

	logger.Log.Info("Request URL", zap.String("url", r.URL.Path))

	// Extract the "id" path parameter from the URL using chi's URLParam method
	idStr := chi.URLParam(r, "id")

	logger.Log.Info("Extracted ID from URL", zap.String("id", idStr))

	if idStr == "" {
		logger.Log.Error("Empty ID parameter")
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Log.Error("Error converting id to integer", zap.String("id", idStr), zap.Error(err))
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	result := conn.First(&product, id)
	if result.Error != nil {
		http.Error(w, "Product not exists", http.StatusNotFound)
		return
	}

	logger.Log.Info("Fetched product", zap.Any("product", product))

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		logger.Log.Error("Error encoding product", zap.Error(err))
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		logger.Log.Error("invalid payload", zap.Error(err))
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}
	if product.Name == "" || product.Price < 0 {
		http.Error(w, "Invalid Product adata ", http.StatusBadRequest)
		return
	}

	conn, err := h.DB.Connect()
	if err != nil {
		logger.Log.Error("DB Connection error", zap.Error(err))
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}

	conn.Create(&product)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&product)

}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	conn, err := h.DB.Connect()
	if err != nil {
		logger.Log.Error("DB Connection error", zap.Error(err))
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid id!", http.StatusBadRequest)
		return
	}
	result := conn.First(&product, id)
	if result.Error != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	//save updaed product
	conn.Save(&product)

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(&product)

}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	conn, err := h.DB.Connect()
	if err != nil {
		logger.Log.Error("DB Connection error", zap.Error(err))
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid prod id", http.StatusNotFound)
		return
	}
	conn.Delete(&models.Product{}, id)
	w.WriteHeader(http.StatusNoContent)

}
