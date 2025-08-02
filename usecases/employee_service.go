package usecases

import (
	"errors"
	"log"
	"net/http"

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

func (svc *EmployeeService) GetAllEmployees() ([]models.Employee, error) {

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
	employees, err := repo.GetAllEmployees()

	if err != nil {
		return nil, err
	}

	return employees, nil
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