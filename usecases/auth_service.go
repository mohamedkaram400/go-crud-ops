package usecases

import (
	"context"
	"errors"
	
	"github.com/google/uuid"
	"github.com/mohamedkaram400/go-crud-ops/helpers"
	"github.com/mohamedkaram400/go-crud-ops/interfaces"
	"github.com/mohamedkaram400/go-crud-ops/models"
	"github.com/mohamedkaram400/go-crud-ops/requests"
	"github.com/mohamedkaram400/go-crud-ops/auth"
)

type AuthService struct {
	Repo interfaces.AuthInterface
}

func NewAuthService(repo interfaces.AuthInterface) *AuthService {
	return &AuthService{Repo: repo}
}

func (svc *AuthService) Register(ctx context.Context, req *requests.RegisterRequest) (*models.Employee, string, error) {
	// Check if username exists
	existing, _ := svc.Repo.GetEmployeeByUsername(ctx, req.UserName)
	if existing != nil {
		return nil, "username already exists", errors.New("username already exists")
	}

	// Hash password
	hashedPwd, err := helpers.HashPassword(req.Password)
	if err != nil {
		return nil, "failed to hash password", err
	}

	// Create employee model
	emp := &models.Employee{
		ID:       	uuid.NewString(),
		Name:       req.Name,
		UserName:   req.UserName,
		Password:   hashedPwd,
		Department: req.Department,
	}

	// Save to DB via repo
	return svc.Repo.Register(emp)
}

func (svc *AuthService) Login(ctx context.Context, req *requests.LoginRequest) (string, error) {
	emp, err := svc.Repo.GetEmployeeByUsername(ctx, req.UserName)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	if err := helpers.CheckPassword(emp.Password, req.Password); err != nil {
		return "", errors.New("invalid username or password")
	}

	// Generate JWT token
	token, err := auth.GenerateJWT(emp.ID, 3)
	if err != nil {
		return "", errors.New("could not generate token")
	}

	return token, nil
}

func (svc *AuthService) Logout(employeeID string) (string, error) {
	return svc.Repo.Logout(employeeID)
}