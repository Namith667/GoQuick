package auth

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/Namith667/GoQuick/internal/logger"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type ContextKey string

const UserRoleKey ContextKey = "userRole"

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}
	return claims, nil
}

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			logger.Log.Warn("Missing authorization token")
			http.Error(w, "Missing authorization token", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ") //trim tokenstring to remove "Bearer" heading

		// Parse token
		claims, err := ParseJWT(tokenString)
		if err != nil {
			logger.Log.Warn("Invalid token", zap.Error(err))
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			logger.Log.Warn("Missing role in token")
			http.Error(w, "Missing role in token", http.StatusUnauthorized)
			return
		}

		// Store role in context
		logger.Log.Info("User authenticated", zap.String("role", role))
		ctx := context.WithValue(r.Context(), UserRoleKey, role)
		logger.Log.Info("Storing role in context", zap.String("role", role))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequireRole(role string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole, ok := r.Context().Value(UserRoleKey).(string)
			logger.Log.Debug("Retrieved user role", zap.String("userRole", userRole))
			logger.Log.Info("Checking role", zap.String("expected", role), zap.String("found", userRole))
			if !ok || userRole != role {
				logger.Log.Info("Insufficient permissions")
				http.Error(w, "Forbidden: Insufficient permissions", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
