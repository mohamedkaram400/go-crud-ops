package handlers


import (
	"context"

	"github.com/mohamedkaram400/go-crud-ops/requests"
	"github.com/mohamedkaram400/go-crud-ops/usecases"
	pb "github.com/mohamedkaram400/go-crud-ops/proto"
)


// AuthGRPCHandler implements the gRPC AuthService
type AuthGRPCHandler struct {
	Service *usecases.AuthService
	pb.UnimplementedAuthServiceServer
}

// gRPC: Register
func (s *AuthGRPCHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	newReq := &requests.RegisterRequest{
		Name:      req.Name,
		UserName:  req.Username,
		Password:  req.Password,
		Department: req.Department,
	}

	emp, msg, err := s.Service.Register(ctx, newReq)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		Message:  msg,
		Employee: &pb.Employee{
			Id:        emp.ID,
			Name:      emp.Name,
			Username:  emp.UserName,
			Department: emp.Department,
		},
	}, nil
}

// gRPC: Login
func (s *AuthGRPCHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	loginReq := &requests.LoginRequest{
		UserName: req.Username,
		Password: req.Password,
	}

	access, refresh, err := s.Service.Login(ctx, loginReq)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

// gRPC: Refresh
func (s *AuthGRPCHandler) Refresh(ctx context.Context, req *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	newAccess, err := s.Service.Refresh(req.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &pb.RefreshResponse{
		AccessToken: newAccess,
	}, nil
}

// gRPC: Logout
func (s *AuthGRPCHandler) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	err := s.Service.Logout(req.EmployeeId)
	if err != nil {
		return nil, err
	}

	return &pb.LogoutResponse{
		Message: "User logged out successfully",
	}, nil
}