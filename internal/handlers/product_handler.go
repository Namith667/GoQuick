package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Namith667/GoQuick/internal/db"
	"github.com/Namith667/GoQuick/internal/models"
	"github.com/gorilla/mux"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	database := db.Connect()

	var products models.Product
	database.Find(&products)

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(products)

}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	database := db.Connect()
	var product models.Product
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	result := database.First(&product, id)
	if result.Error != nil {
		http.Error(w, "Product not exists", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	database := db.Connect()
	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	database.Create(&product)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&product)

}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	database := db.Connect()
	var product models.Product
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid id!", http.StatusBadRequest)
		return
	}
	result := database.First(&product, id)
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
	database.Save(&product)

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(&product)

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	database := db.Connect()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid prod id", http.StatusNotFound)
		return
	}
	database.Delete(&models.Product{}, id)
	w.WriteHeader(http.StatusNoContent)

}
