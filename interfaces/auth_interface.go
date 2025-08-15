package interfaces

import (
	"context"

	"github.com/mohamedkaram400/go-crud-ops/models"
)

type AuthInterface interface {
    GetEmployeeByUsername(ctx context.Context, userName string) (*models.Employee, error)
    Register(emp *models.Employee) (*models.Employee, string, error)
    Logout(employeeID string) (string, error)
}

