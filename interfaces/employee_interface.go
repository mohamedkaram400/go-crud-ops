package interfaces

import (
	"github.com/mohamedkaram400/go-crud-ops/models"
)

type EmployeeInterface interface {
    InsertEmployee(emp *models.Employee) (*models.Employee, error)
    FindEmployeeByID(employeeID string) (*models.Employee, error)
    GetAllEmployees(skip int, limit int) ([]*models.Employee, int, error)
    UpdateEmployee(employeeID string, newEmployee *models.Employee) (int, error)
    DeleteEmployee(employeeID string) (int, error)
}