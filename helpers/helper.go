package helpers

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/mohamedkaram400/go-crud-ops/models"
)

type EmployeeDTO struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	UserName   string `json:"username"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(hashed, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
}

func ConvertEmployeesToDTOs(employees []*models.Employee) []*EmployeeDTO {
	dtos := make([]*EmployeeDTO, len(employees))
	for i, emp := range employees {
		dtos[i] = ConvertToEmployeeDTO(emp)
	}
	return dtos
}

func ConvertToEmployeeDTO(emp *models.Employee) *EmployeeDTO {
	return &EmployeeDTO{
		ID:         emp.ID,
		Name:       emp.Name,
		Department: emp.Department,
		UserName:   emp.UserName,
	}
}

