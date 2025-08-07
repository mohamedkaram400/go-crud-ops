package usecases

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mohamedkaram400/go-crud-ops/models"
	"github.com/mohamedkaram400/go-crud-ops/repository"
	"github.com/mohamedkaram400/go-crud-ops/requests"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService struct {
	MongoCollection *mongo.Collection
}

func (svc *EmployeeService) CreateEmployee(employee *models.Employee) (*models.Employee, error) {

	employee.EmployeeID = uuid.NewString()

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
	_, err := repo.InsertEmployee(employee)
	
	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (svc *EmployeeService) GetAllEmployees(pageStr, limitStr string) ([]models.Employee, int, int, int, error) {

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

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	employees, totalCount, err := repo.GetAllEmployees(skip, limit)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	return employees, totalCount, page, limit, nil
}

func (svc *EmployeeService) FindEmployeeByID(r *http.Request) (*models.Employee, error) {

	empID := mux.Vars(r)["uuid"]
	log.Println("employee id", empID)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
	employee, err := repo.FindEmployeeByID(empID)

	if err != nil {
		return nil, err
	}
	
	return employee, nil
}

func (svc *EmployeeService) UpdateEmployee(r *http.Request, reqData *requests.UpdateEmployeeRequest) (int, error) {

	// Step 2: Get employee ID from path
	vars := mux.Vars(r)
	employeeID := vars["uuid"]
	if employeeID == "" {
		return 0, errors.New("Employee ID is required in path")
	}

	// Convert request to model
	employee := &models.Employee{
		EmployeeID: employeeID,
		Name:       reqData.Name,
		Department: reqData.Department,
	}

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
	count, err := repo.UpdateEmployee(employeeID, employee)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (svc *EmployeeService) DeleteEmployee(r *http.Request) (int, error) {
	
	employeeID := mux.Vars(r)["uuid"]
	log.Println("employee id", employeeID)

	if employeeID == "" {
		return 0, errors.New("invalid employee id")
	}

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
	count, err := repo.DeleteEmployee(employeeID)
	if err != nil {
		return 0, err
	}

	return count, nil
}