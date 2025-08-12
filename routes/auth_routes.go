package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohamedkaram400/go-crud-ops/handlers"
)

func RegisterAuthRoutes(router *mux.Router, authHandler *handlers.AuthHandler) {
	// Health Check
	router.HandleFunc("/health", HealthHandler).Methods(http.MethodGet)

	// Auth Routes
	router.HandleFunc("/register", authHandler.RegisterHandler)
	router.HandleFunc("/login", authHandler.LoginHandler)
	router.HandleFunc("/logout", authHandler.LogoutHandler)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([] byte("API v1 up"))
}		
