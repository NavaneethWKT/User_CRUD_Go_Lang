package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/NavaneethWKT/CRUD-Go-lang/internal/jwt"
)

type ContextKey string

const (
	UserIDKey ContextKey = "user_id"
	UserEmailKey ContextKey = "user_email"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Authorization header is required")
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Invalid authorization header format. Use: Bearer <token>")
			return
		}
		tokenString := parts[1]
		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Invalid or expired token: " + err.Error())
			return
		}
		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		ctx = context.WithValue(ctx, UserEmailKey, claims.Email)
		r = r.WithContext(ctx)
		next(w, r)
	}
}

func GetUserID(ctx context.Context) string {
	if userID, ok := ctx.Value(UserIDKey).(string); ok {
		return userID
	}
	return ""
}

func GetUserEmail(ctx context.Context) string {
	if email, ok := ctx.Value(UserEmailKey).(string); ok {
		return email
	}
	return ""
}

