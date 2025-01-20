package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Namith667/GoQuick/internal/services"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (ah *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	user, err := ah.AuthService.RegisterUser(input.Username, input.Email, input.Password)
	if err != nil {
		http.Error(w, "User registration failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if error := json.NewDecoder(r.Body).Decode(&input); error != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	token, err := ah.AuthService.AuthenticateUser(input.Email, input.Password)
	if err != nil {
		http.Error(w, "invalid credentials from token gen", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
