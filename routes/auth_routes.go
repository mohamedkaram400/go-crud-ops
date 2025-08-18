package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohamedkaram400/go-crud-ops/handlers"
	"github.com/mohamedkaram400/go-crud-ops/middlewares"
)

func RegisterAuthRoutes(router *mux.Router, authHandler *handlers.AuthHandler) {
	// Health Check
	router.HandleFunc("/health", HealthHandler).Methods(http.MethodGet)

	// Public routes
	router.HandleFunc("/register", authHandler.RegisterHandler).Methods(http.MethodPost)
	router.HandleFunc("/login", authHandler.LoginHandler).Methods(http.MethodPost)

	router.HandleFunc("/refresh", authHandler.RefreshHandler).Methods(http.MethodPost)

	// Protected routes (JWT required)
	auth := router.PathPrefix("").Subrouter()
	auth.Use(middlewares.JWTAuth)
	auth.HandleFunc("/logout", authHandler.LogoutHandler).Methods(http.MethodPost)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([] byte("API v1 up"))
}		



// router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "The "+r.Method+" method is not supported for this route. Supported methods: POST.", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// Call your handler here
// 	authHandler.LogoutHandler(w, r)
// })