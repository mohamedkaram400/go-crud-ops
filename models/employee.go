package models

type Employee struct {
	EmployeeID string `json:"employeeId"`
	Name       string `json:"name"`
	Department string `json:"department"`
}