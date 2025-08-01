package usecases

import (
	// "encoding/json"
	// "log"
	// "net/http"

	"github.com/google/uuid"
	// "github.com/gorilla/mux"
	"github.com/mohamedkaram400/go-crud-ops/models"
	"github.com/mohamedkaram400/go-crud-ops/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService struct {
	MongoCollection *mongo.Collection
}

// type Response struct {
// 	Data  interface{} `json:"data"`
// 	Error string      `json:"error,omitempty"`
// }

func (svc *EmployeeService) CreateEmployee(emp *models.Employee) (string, error) {

	// Business logic
	emp.EmployeeID = uuid.NewString()

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
	_, err := repo.InsertEmployee(emp)
	
	if err != nil {
		return "", err
	}

	return emp.EmployeeID, nil
}

// func (svc *EmployeeService) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Content-Type", "application/json")

// 	res := &Response{}
// 	defer json.NewEncoder(w).Encode(res)

// 	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

// 	// Insert employee
// 	emp, err := repo.GetAllEmployees()

// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		log.Print("error", err)
// 		res.Error = err.Error()
// 		return
// 	}
	
// 	res.Data = emp
// 	w.WriteHeader(http.StatusOK)
// }

// func (svc *EmployeeService) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Content-Type", "application/json")

// 	res := &Response{}
// 	defer json.NewEncoder(w).Encode(res)

// 	empID := mux.Vars(r)["id"]
// 	log.Println("employee id", empID)

// 	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

// 	// Insert employee
// 	emp, err := repo.FindEmployeeByID(empID)

// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		log.Print("error", err)
// 		res.Error = err.Error()
// 		return
// 	}
	
// 	res.Data = emp
// 	w.WriteHeader(http.StatusOK)
// }

// func (svc *EmployeeService) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Content-Type", "application/json")
	
// 	res := &Response{}
// 	defer json.NewEncoder(w).Encode(res)

// 	employeeID := mux.Vars(r)["id"]
// 	log.Println("employee id", employeeID)

// 	if  employeeID == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		log.Print("Invalid employee id")
// 		res.Error = "Invalid employee id"
// 		return
// 	}
	
// 	var emp models.Employee

// 	err := json.NewDecoder(r.Body).Decode(&emp)

// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		log.Print("Invalid body", err)
// 		res.Error = err.Error()
// 		return
// 	}

// 	emp.EmployeeID = employeeID

// 	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

// 	// Update employee
// 	count, err := repo.UpdateEmployee(employeeID, &emp)


// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		log.Print("Invalid body", err)
// 		res.Error = err.Error()
// 		return
// 	}

// 	res.Data = count
// 	w.WriteHeader(http.StatusOK)

// 	log.Println("Employee inserted with id", employeeID, emp)

// }

// func (svc *EmployeeService) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Content-Type", "application/json")

// 	res := &Response{}
// 	defer json.NewEncoder(w).Encode(res)

// 	employeeID := mux.Vars(r)["id"]
// 	log.Println("employee id", employeeID)

// 	if  employeeID == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		log.Print("Invalid employee id")
// 		res.Error = "Invalid employee id"
// 		return
// 	}

// 	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

// 	count, err := repo.DeleteEmployee(employeeID)

// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		log.Print("Invalid ", err)
// 		res.Error = err.Error()
// 		return
// 	}

// 	res.Data = count
// 	w.WriteHeader(http.StatusOK)
// }