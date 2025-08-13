package usecases

import (
	"context"
	"errors"
	"time"
	
	"github.com/mohamedkaram400/go-crud-ops/helpers"
	"github.com/mohamedkaram400/go-crud-ops/interfaces"
	// "github.com/mohamedkaram400/go-crud-ops/models"
	"github.com/mohamedkaram400/go-crud-ops/requests"
	"github.com/mohamedkaram400/go-crud-ops/auth"
)

type AuthService struct {
	Repo interfaces.AuthInterface
}

func NewAuthService(repo interfaces.AuthInterface) *AuthService {
	return &AuthService{Repo: repo}
}

// func (svc *AuthService) register() (*models.Employee, error) {

// }


func (svc *AuthService) Login(ctx context.Context, req *requests.LoginRequest) (string, error) {
	emp, err := svc.Repo.GetEmployeeByUsername(ctx, req.UserName)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	if err := helpers.CheckPassword(emp.Password, req.Password); err != nil {
		return "", errors.New("invalid username or password")
	}

	// Generate JWT token
	token, err := auth.GenerateJWT(emp.ID, time.Hour*1)
	if err != nil {
		return "", errors.New("could not generate token")
	}

	return token, nil
}

// func (svc *AuthService) Logout(employeeID string) error {

// }