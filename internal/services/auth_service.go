package services

import (
	"errors"
	"os"
	"time"

	"github.com/Namith667/GoQuick/internal/config"
	"github.com/Namith667/GoQuick/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{DB: db}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func generateJWT(user models.User) (string, error) {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		return "", errors.New("JWT SECRET DOES NOT EXIST")
	}

	expTime := config.GetExpirationTime()
	claims := jwt.MapClaims{
		"username": user.Name,
		"email":    user.Email,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * time.Duration(expTime)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func (as *AuthService) RegisterUser(username, email, password string) (models.User, error) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Name:     username,
		Password: hashedPassword,
		Email:    email,
		Role:     "user",
	}

	result := as.DB.Create(&user)
	return user, result.Error
}

func (as *AuthService) AuthenticateUser(email, password string) (string, error) {
	var user models.User
	result := as.DB.Where("email=?", email).First(&user)
	if result.Error != nil {
		return "", errors.New("invalid credentials")
	}
	if !VerifyPassword(user.Password, password) {
		return "", errors.New("invalid password")
	}
	token, err := generateJWT(user)
	if err != nil {
		return "", err
	}
	return token, nil

}
