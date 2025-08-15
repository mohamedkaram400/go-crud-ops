package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohamedkaram400/go-crud-ops/handlers"
)

func RegisterEmployeeRoutes(router *mux.Router, employeeHandler *handlers.EmployeeHandler) {

	// Employee Routes
	router.HandleFunc("/create", employeeHandler.CreateEmployee).Methods(http.MethodPost)
	router.HandleFunc("/all", 			employeeHandler.GetAllEmployees).Methods(http.MethodGet)
	router.HandleFunc("/show/{uuid}", 	employeeHandler.GetEmployeeByID).Methods(http.MethodGet)
	router.HandleFunc("/edit/{uuid}", 	employeeHandler.UpdateEmployee).Methods(http.MethodPut)
	router.HandleFunc("/delete/{uuid}", 	employeeHandler.DeleteEmployee).Methods(http.MethodDelete)
}

