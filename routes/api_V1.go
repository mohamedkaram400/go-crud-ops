package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohamedkaram400/go-crud-ops/handlers"
)

func RegisterAPIV1Routes(router *mux.Router, employeeHandler *handlers.EmployeeHandler, authHandler *handlers.AuthHandler) {
	api := router.PathPrefix("/api/v1").Subrouter()

	// Health Check
	api.HandleFunc("/health", HealthHandler).Methods(http.MethodGet)

	// Auth Routes
	api.HandleFunc("/register", authHandler.RegisterHandler)
	api.HandleFunc("/login", authHandler.LoginHandler)
	api.HandleFunc("/logout", authHandler.LogoutHandler)

	// Employee Routes
	api.HandleFunc("/employees/create", employeeHandler.CreateEmployee).Methods(http.MethodPost)
	api.HandleFunc("/employees/all", 			employeeHandler.GetAllEmployees).Methods(http.MethodGet)
	api.HandleFunc("/employees/show/{uuid}", 	employeeHandler.GetEmployeeByID).Methods(http.MethodGet)
	api.HandleFunc("/employees/edit/{uuid}", 	employeeHandler.UpdateEmployee).Methods(http.MethodPut)
	api.HandleFunc("/employees/delete/{uuid}", 	employeeHandler.DeleteEmployee).Methods(http.MethodDelete)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([] byte("API v1 up"))

}		