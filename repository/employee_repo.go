package repository

import (
	"context"
	"fmt"
	"log"
	"errors"

	"github.com/mohamedkaram400/go-crud-ops/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepo struct {
	MongoCollection *mongo.Collection
}

func (r *EmployeeRepo) InsertEmployee(emp *models.Employee) (*models.Employee, error) {
	_, err := r.MongoCollection.InsertOne(context.Background(), emp)

	if err != nil {
		return nil, err
	}

	return emp, nil
}

func (r *EmployeeRepo) FindEmployeeByID(employeeID string) (*models.Employee, error) {
	var emp models.Employee

	err := r.MongoCollection.FindOne(context.Background(),
		bson.D{{Key: "employeeid", Value: employeeID}}).Decode(&emp)

	if err != nil {
		return nil, err
	}

	return &emp, nil
}

func (r *EmployeeRepo) GetAllEmployees() ([]models.Employee, error) {
	result, err := r.MongoCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	var emps []models.Employee
	err = result.All(context.Background(), &emps)

	if err != nil {
		return nil, fmt.Errorf("results decode error %s", err.Error())
	}

	return emps, nil
}

func (r *EmployeeRepo) UpdateEmployee(employeeID string, newEmployee *models.Employee) (int, error) {
	update := bson.D{}

	if newEmployee.Name != "" {
		update = append(update, bson.E{Key: "name", Value: newEmployee.Name})
	}

	if newEmployee.Department != "" {
		update = append(update, bson.E{Key: "department", Value: newEmployee.Department})
	}

	if len(update) == 0 {
		return 0, errors.New("no fields to update")
	}


	result, err := r.MongoCollection.UpdateOne(context.Background(),
		bson.D{{Key: "employeeid", Value: employeeID}},
		bson.D{{Key: "$set", Value: update}})

	if err != nil {
		return 0, err
	}

	return int(result.ModifiedCount), nil
}

func (r *EmployeeRepo) DeleteEmployee(employeeID string) (int, error) {

	log.Println("employee id", employeeID)

	result, err := r.MongoCollection.DeleteOne(context.Background(),
		bson.D{{Key: "employeeid", Value: employeeID}})

	if err != nil {
		return 0, err
	}

	log.Printf("Deleting employee with ID: %s\n", employeeID)

	return int(result.DeletedCount), nil
}
