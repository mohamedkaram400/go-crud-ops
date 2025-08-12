package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohamedkaram400/go-crud-ops/handlers"
)

func RegisterEmployeeRoutes(router *mux.Router, employeeHandler *handlers.EmployeeHandler) {

	// Employee Routes
	router.HandleFunc("/employees/create", employeeHandler.CreateEmployee).Methods(http.MethodPost)
	router.HandleFunc("/employees/all", 			employeeHandler.GetAllEmployees).Methods(http.MethodGet)
	router.HandleFunc("/employees/show/{uuid}", 	employeeHandler.GetEmployeeByID).Methods(http.MethodGet)
	router.HandleFunc("/employees/edit/{uuid}", 	employeeHandler.UpdateEmployee).Methods(http.MethodPut)
	router.HandleFunc("/employees/delete/{uuid}", 	employeeHandler.DeleteEmployee).Methods(http.MethodDelete)
}

