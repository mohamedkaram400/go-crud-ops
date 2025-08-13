package usecases

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
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

func (svc *EmployeeService) FindEmployeeByID(r *http.Request) (*models.Employee, error) {

	empID := mux.Vars(r)["uuid"]
	log.Println("employee id", empID)

	employee, err := svc.Repo.FindEmployeeByID(empID)

	if err != nil {
		return nil, err
	}
	
	return employee, nil
}

func (svc *EmployeeService) UpdateEmployee(r *http.Request, reqData *requests.UpdateEmployeeRequest) (int, error) {

	// Step 2: Get employee ID from path
	vars := mux.Vars(r)
	id := vars["uuid"]
	if id == "" {
		return 0, errors.New("Employee ID is required in path")
	}

	// Convert request to model
	employee := &models.Employee{
		ID: id,
		Name:       reqData.Name,
		Department: reqData.Department,
	}

	count, err := svc.Repo.UpdateEmployee(id, employee)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (svc *EmployeeService) DeleteEmployee(r *http.Request) (int, error) {
	
	id := mux.Vars(r)["uuid"]
	log.Println("employee id", id)

	if id == "" {
		return 0, errors.New("invalid employee id")
	}

	count, err := svc.Repo.DeleteEmployee(id)
	if err != nil {
		return 0, err
	}

	return count, nil
}