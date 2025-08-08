package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mohamedkaram400/go-crud-ops/helpers"
	"github.com/mohamedkaram400/go-crud-ops/models"
	"github.com/mohamedkaram400/go-crud-ops/requests"
	"github.com/mohamedkaram400/go-crud-ops/usecases"
)

type PaginatedResult struct {
	Message    string              `json:"message,omitempty"`
	Error      string              `json:"error,omitempty"`
	Data       []*helpers.EmployeeDTO   `json:"data,omitempty"`
	TotalCount int                 `json:"totalCount"`
	Page       int                 `json:"page"`
	Limit      int                 `json:"limit"`
}

type EmployeeResponse struct {
	Message     string 			`json:"message,omitempty"`
	Data  		interface{} 	`json:"data,omitempty"`
	Error 		string      	`json:"error,omitempty"`
}

type EmployeeHandler struct {
	Service *usecases.EmployeeService
}

func (h *EmployeeHandler) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	fmt.Println(pageStr, limitStr)

	employees, message, totalCount, page, limit, err := h.Service.GetAllEmployees(pageStr, limitStr)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&PaginatedResult{
			Error: err.Error(),
		})
		return
	}

	res := &PaginatedResult{
		Message:  message,
		Data: helpers.ConvertEmployeesToDTOs(employees),
		Page: page,
		Limit: limit,
		TotalCount: totalCount,
	}

	defer json.NewEncoder(w).Encode(res)

	w.WriteHeader(http.StatusOK)
}

func (h *EmployeeHandler) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &EmployeeResponse{}
	defer json.NewEncoder(w).Encode(res)

	// Get employee
	employee, err := h.Service.FindEmployeeByID(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("error", err)
		res.Error = err.Error()
		return
	}
	
	res.Message = "Employee returned successfully"
	res.Data = employee
	w.WriteHeader(http.StatusOK)
}

func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := &EmployeeResponse{}
	defer json.NewEncoder(w).Encode(res)

	fmt.Println(r)
	// Step 1: Validate
	req, err := requests.ParseAndValidateCreateEmployee(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = err.Error()
		return
	}

	emp := &models.Employee{
		Name:       req.Name,
		Department: req.Department,
		UserName: req.UserName,
		Password: req.Password,
	}

	// Step 2: Call service
	employee, err := h.Service.CreateEmployee(emp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		return
	}

	// Step 3: Return response
	res.Message = "Employee created successfully"
	res.Data = employee
	w.WriteHeader(http.StatusCreated)
}

func (h *EmployeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := &EmployeeResponse{}
	defer json.NewEncoder(w).Encode(res)


	// Step 1: Validate
	req, err := requests.ParseAndValidateUpdateEmployee(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = err.Error()
		return
	}

	// Step 2: Call service
	count, err := h.Service.UpdateEmployee(r, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		return
	}

	// Step 3: Return response
	res.Message = "Employee updated successfully"
	res.Data = count
	w.WriteHeader(http.StatusOK)
}

func (h *EmployeeHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &EmployeeResponse{}
	defer json.NewEncoder(w).Encode(res)

	count, err := h.Service.DeleteEmployee(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Invalid ", err)
		res.Error = err.Error()
		return
	}

	res.Message = "Employee deleted successfully"
	res.Data = count
	w.WriteHeader(http.StatusOK)
}