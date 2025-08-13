package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/mohamedkaram400/go-crud-ops/auth"
)

type contextKey string

const EmployeeIDKey contextKey = "employeeID"

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get token from header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		// Validate token
		employeeID, err := auth.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Store employee ID in context for later use
		ctx := context.WithValue(r.Context(), EmployeeIDKey, employeeID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
