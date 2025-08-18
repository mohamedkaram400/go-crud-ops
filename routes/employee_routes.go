package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohamedkaram400/go-crud-ops/handlers"
	"github.com/mohamedkaram400/go-crud-ops/middlewares"
)


func RegisterEmployeeRoutes(router *mux.Router, employeeHandler *handlers.EmployeeHandler) {
	// Protected employees group
	employees := router.PathPrefix("/employees").Subrouter()
	employees.Use(middlewares.JWTAuth)

	employees.HandleFunc("/create", employeeHandler.CreateEmployee).Methods(http.MethodPost)
	employees.HandleFunc("/all", employeeHandler.GetAllEmployees).Methods(http.MethodGet)
	employees.HandleFunc("/show/{uuid}", employeeHandler.GetEmployeeByID).Methods(http.MethodGet)
	employees.HandleFunc("/edit/{uuid}", employeeHandler.UpdateEmployee).Methods(http.MethodPut)
	employees.HandleFunc("/delete/{uuid}", employeeHandler.DeleteEmployee).Methods(http.MethodDelete)
}