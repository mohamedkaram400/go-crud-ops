package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/mohamedkaram400/go-crud-ops/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepo struct {
	MongoCollection *mongo.Collection
}

func (r *EmployeeRepo) InsertEmployee(emp *models.Employee) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), emp)

	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (r *EmployeeRepo) FindEmployeeByID(employeeID string) (*models.Employee, error) {
	var emp models.Employee

	err := r.MongoCollection.FindOne(context.Background(),
		bson.D{{Key: "employeeID", Value: employeeID}}).Decode(&emp)

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
	result, err := r.MongoCollection.UpdateOne(context.Background(),
		bson.D{{Key: "employeeID", Value: employeeID}},
		bson.D{{Key: "$set", Value: newEmployee}})

	if err != nil {
		return 0, err
	}

	return int(result.ModifiedCount), nil
}

func (r *EmployeeRepo) DeleteEmployee(employeeID string) (int, error) {
	result, err := r.MongoCollection.DeleteOne(context.Background(),
		bson.D{{Key: "EmployeeID", Value: employeeID}})

	if err != nil {
		return 0, err
	}

	log.Printf("Deleting employee with ID: %s\n", employeeID)

	return int(result.DeletedCount), nil
}
