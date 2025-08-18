package usecases

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mohamedkaram400/go-crud-ops/helpers"
	"github.com/mohamedkaram400/go-crud-ops/interfaces"
	"github.com/mohamedkaram400/go-crud-ops/models"
	"github.com/mohamedkaram400/go-crud-ops/requests"
	"github.com/mohamedkaram400/go-crud-ops/auth"
	"github.com/mohamedkaram400/go-crud-ops/internal/redis"
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

func (svc *AuthService) Login(ctx context.Context, req *requests.LoginRequest) (string, string, error) {
	emp, err := svc.Repo.GetEmployeeByUsername(ctx, req.UserName)
	if err != nil {
		return "", "", errors.New("invalid username or password")
	}

	if err := helpers.CheckPassword(emp.Password, req.Password); err != nil {
		return "", "", errors.New("invalid username or password")
	}

	// Access token (short-lived, 15 min)
	accessToken, err := auth.GenerateAccessToken(emp.ID, 1) // 1 hour for now
	if err != nil {
		return "", "", errors.New("could not generate access token")
	}

	// Refresh token (long-lived, 7 days)
	refreshToken, err := auth.GenerateRefreshToken(emp.ID, 7)
	if err != nil {
		return "", "", errors.New("could not generate refresh token")
	}

	// Store refresh token in Redis or DB
	err = redisclient.Client.Set(ctx, emp.ID, refreshToken, 7*24*time.Hour).Err()
	if err != nil {
		return "", "", errors.New("failed to store refresh token")
	}

	return accessToken, refreshToken, nil
}

func (svc *AuthService) Refresh(refreshToken string) (string, error) {
	// 1. Validate refresh token
	employeeID, err := auth.ValidateJWT(refreshToken)
	if err != nil {
		return "", fmt.Errorf("invalid refresh token: %w", err)
	}

	// 2. Check if refresh token exists in Redis (not revoked)
	storedToken, err := redisclient.Client.Get(context.Background(), employeeID).Result()
	if err != nil || storedToken != refreshToken {
		return "", fmt.Errorf("refresh token not valid or revoked")
	}

	// 3. Generate new access token
	newAccessToken, err := auth.GenerateAccessToken(employeeID, int(time.Hour.Seconds())) // expires in 1h
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}

func (svc *AuthService) Logout(employeeID string) error {
	return redisclient.Client.Del(context.Background(), employeeID).Err()
}
