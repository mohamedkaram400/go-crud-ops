package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohamedkaram400/go-crud-ops/handlers"
)

func RegisterAPIV1Routes(router *mux.Router, empHandler *handlers.EmployeeHandler) {
	api := router.PathPrefix("/api/v1/employees").Subrouter()

	// Health Check
	api.HandleFunc("/health", HealthHandler).Methods(http.MethodGet)

	// Employee Routes
	api.HandleFunc("/create", empHandler.CreateEmployee).Methods(http.MethodPost)
	// api.HandleFunc("/all", 			empHandler.GetAllEmployees).Methods(http.MethodGet)
	// api.HandleFunc("/show/{id}", 	empHandler.GetEmployeeByID).Methods(http.MethodGet)
	// api.HandleFunc("/edit/{id}", 	empHandler.UpdateEmployee).Methods(http.MethodPut)
	// api.HandleFunc("/delete/{id}", 	empHandler.DeleteEmployee).Methods(http.MethodDelete)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([] byte("API v1 up"))

}		