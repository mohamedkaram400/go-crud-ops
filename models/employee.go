package models

type Employee struct {
	EmployeeID string `json:"employee_id,omitempty"`
	Name       string `json:"name"`
	Department string `json:"department"`
}