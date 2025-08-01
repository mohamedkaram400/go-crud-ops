package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mohamedkaram400/go-crud-ops/requests"
	"github.com/mohamedkaram400/go-crud-ops/usecases"
	"github.com/mohamedkaram400/go-crud-ops/models"
)

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

type EmployeeHandler struct {
	Service *usecases.EmployeeService
}

func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := &Response{}
	
	defer json.NewEncoder(w).Encode(res)

	
	// Step 1: Validate
	req, err := requests.ParseAndValidateEmployee(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = err.Error()
		return
	}

	emp := &models.Employee{
		Name:       req.Name,
		Department: req.Department,
	}

	// Step 2: Usecase
	id, err := h.Service.CreateEmployee(emp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		return
	}

	// Step 3: Return response
	res.Data = map[string]string{"employee_id": id}
}