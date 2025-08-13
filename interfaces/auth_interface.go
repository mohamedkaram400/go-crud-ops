package interfaces

import (
	"context"

	"github.com/mohamedkaram400/go-crud-ops/models"
)

type AuthInterface interface {
    GetEmployeeByUsername(ctx context.Context, userName string) (models.Employee, error)
    // Register(emp *models.Employee) (*models.Employee, error)
    // Login(emp *models.Employee) (*models.Employee, error)
    // Logout(emp *models.Employee) (*models.Employee, error)
}

