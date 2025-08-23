package usecases

import (
	"errors"
	"log"
	"strconv"

	"github.com/google/uuid"
	"github.com/mohamedkaram400/go-crud-ops/helpers"
	"github.com/mohamedkaram400/go-crud-ops/interfaces"
	"github.com/mohamedkaram400/go-crud-ops/models"
	"github.com/mohamedkaram400/go-crud-ops/repository"
	"github.com/mohamedkaram400/go-crud-ops/requests"
)

type EmployeeService struct {
    Repo interfaces.EmployeeInterface
}

func NewEmployeeService(repo *repository.EmployeeRepo) *EmployeeService {
    return &EmployeeService{repo}
}

func (svc *EmployeeService) CreateEmployee(employee *models.Employee) (*models.Employee, error) {

	hashedPassword, err := helpers.HashPassword(employee.Password)
	if err != nil {
		return nil, err
	}

	employee.ID = uuid.NewString()
	employee.Password = hashedPassword

	_, err = svc.Repo.InsertEmployee(employee)
	
	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (svc *EmployeeService) GetAllEmployees(pageStr, limitStr string) ([]*models.Employee, string, int, int, int, error) {

	page := 1
	limit := 10

	var err error
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}
	}

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit < 1 {
			limit = 10
		}
	}

	skip := (page - 1) * limit

	employees, totalCount, err := svc.Repo.GetAllEmployees(skip, limit)
	if err != nil {
		return nil, "", 0, 0, 0, err
	}

	if employees == nil {
		return nil, "Not found data", 0, 0, 0, nil
	}
	return employees, "Employees returned successfully", totalCount, page, limit, nil
}

func (svc *EmployeeService) FindEmployeeByID(employeeID string) (*models.Employee, error) {

	log.Println("employee id", employeeID)

	employee, err := svc.Repo.FindEmployeeByID(employeeID)

	if err != nil {
		return nil, err
	}
	
	return employee, nil
}

func (svc *EmployeeService) UpdateEmployeeByID(employeeID string, reqData *requests.UpdateEmployeeRequest) (int, error) {

	// Step 2: Get employee ID from path
	if employeeID == "" {
		return 0, errors.New("employee ID is required in path")
	}

	// Convert request to model
	employee := &models.Employee{
		ID: employeeID,
		Name:       reqData.Name,
		Department: reqData.Department,
	}

	count, err := svc.Repo.UpdateEmployee(employeeID, employee)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (svc *EmployeeService) DeleteEmployeeByID(employeeID string) (int, error) {
	
	log.Println("employee id", employeeID)

	if employeeID == "" {
		return 0, errors.New("invalid employee id")
	}

	count, err := svc.Repo.DeleteEmployee(employeeID)
	if err != nil {
		return 0, err
	}

	return count, nil
}