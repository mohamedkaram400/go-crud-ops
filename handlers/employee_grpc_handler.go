package handlers

import (
	"context"
	"strconv"

	pb "github.com/mohamedkaram400/go-crud-ops/proto"
	"github.com/mohamedkaram400/go-crud-ops/models"
	"github.com/mohamedkaram400/go-crud-ops/usecases"
	"github.com/mohamedkaram400/go-crud-ops/requests"
)

type EmployeeGRPCHandler struct {
	Service *usecases.EmployeeService
	pb.UnimplementedEmployeeServiceServer
}

// GetAllEmployees
func (s *EmployeeGRPCHandler) GetAllEmployees(ctx context.Context, req *pb.GetAllEmployeesRequest) (*pb.EmployeeListResponse, error) {
	page := strconv.Itoa(int(req.Page))
	limit := strconv.Itoa(int(req.Limit))

	employees, message, totalCount, pageInt, limitInt, err := s.Service.GetAllEmployees(page, limit)
	if err != nil {
		return nil, err
	}

	// Convert to protobuf employees
	var pbEmployees []*pb.Employee
	for _, emp := range employees {
		pbEmployees = append(pbEmployees, &pb.Employee{
			Id:        emp.ID,
			Name:      emp.Name,
			Username:  emp.UserName,
			Department: emp.Department,
		})
	}

	return &pb.EmployeeListResponse{
		Message:    message,
		Employees:       pbEmployees,
		TotalCount: int32(totalCount),
		Page:       int32(pageInt),
		Limit:      int32(limitInt),
	}, nil
}

// GetEmployeeByID
func (s *EmployeeGRPCHandler) GetEmployeeByID(ctx context.Context, req *pb.GetEmployeeRequest) (*pb.EmployeeResponse, error) {
	emp, err := s.Service.FindEmployeeByID(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.EmployeeResponse{
		Message: "Employee returned successfully",
		Employee: &pb.Employee{
			Id:        emp.ID,
			Name:      emp.Name,
			Username:  emp.UserName,
			Department: emp.Department,
		},
	}, nil
}

// CreateEmployee
func (s *EmployeeGRPCHandler) CreateEmployee(ctx context.Context, req *pb.CreateEmployeeRequest) (*pb.EmployeeResponse, error) {
	emp := &models.Employee{
		Name:       req.Name,
		UserName:   req.Username,
		Department: req.Department,
		Password:   req.Password,
	}

	newEmp, err := s.Service.CreateEmployee(emp)
	if err != nil {
		return nil, err
	}

	return &pb.EmployeeResponse{
		Message: "Employee created successfully",
		Employee: &pb.Employee{
			Id:        newEmp.ID,
			Name:      newEmp.Name,
			Username:  newEmp.UserName,
			Department: newEmp.Department,
		},
	}, nil
}

// UpdateEmployee
func (s *EmployeeGRPCHandler) UpdateEmployee(ctx context.Context, req *pb.UpdateEmployeeRequest) (*pb.EmployeeResponse, error) {
	reqData := &requests.UpdateEmployeeRequest{ // ✅ wrap fields into service request
		Name:       req.Name,
		UserName:   req.Username,
		Password:   req.Password,
		Department: req.Department,
	}

	_, err := s.Service.UpdateEmployeeByID(req.Id, reqData)
	if err != nil {
		return nil, err
	}

	return &pb.EmployeeResponse{
		Message: "Employee updated successfully",
		Employee: &pb.Employee{
			Id:        req.Id,
			Name:      req.Name,
			Username:  req.Username,
			Department: req.Department,
		},
	}, nil
}

// DeleteEmployee
func (s *EmployeeGRPCHandler) DeleteEmployee(ctx context.Context, req *pb.DeleteEmployeeRequest) (*pb.EmployeeResponse, error) {
	_, err := s.Service.DeleteEmployeeByID(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.EmployeeResponse{
		Message: "Employee deleted successfully",
	}, nil
}