package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Namith667/GoQuick/internal/logger"
	"github.com/Namith667/GoQuick/internal/middleware/auth"
	"github.com/Namith667/GoQuick/internal/models"
	"go.uber.org/zap"
)

type AuthHandler struct {
	AuthService *auth.AuthService
}

func NewAuthHandler(authService *auth.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (ah *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var input models.RegisterUserInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logger.Log.Error("Invalid user input", zap.Error(err))
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	user, err := ah.AuthService.RegisterUser(input.Username, input.Email, input.Password, input.Role)
	if err != nil {
		logger.Log.Error("User registration failed", zap.Error(err))
		http.Error(w, "User registration failed", http.StatusInternalServerError)
		return
	}
	//added a reg-user response(see models package) to avoid sending back password
	response := models.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}
	logger.Log.Info("User Registered successfully", zap.String("name", input.Username))
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input models.LoginUserInput
	if error := json.NewDecoder(r.Body).Decode(&input); error != nil {
		logger.Log.Error("Invalid user input", zap.Error(error))
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	token, err := ah.AuthService.AuthenticateUser(input.Email, input.Password)
	if err != nil {
		logger.Log.Error("Invalid Credentials", zap.Error(err))
		http.Error(w, "invalid credentials from token gen", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
