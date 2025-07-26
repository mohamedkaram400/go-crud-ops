package repository

import (
	"context"
	"fmt"

	"github.com/mohamedkaram400/go-crud-ops/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepo struct {
	MongoCollection *mongo.Collection
}

func (r *EmployeeRepo) InsertEmployee(emp *model.Employee) (*mongo.InsertOneResult, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), emp)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *EmployeeRepo) FindEmployeeByID(employeeID string) (*model.Employee, error) {
	var emp model.Employee

	err := r.MongoCollection.FindOne(context.Background(),
		bson.D{{Key: "employeeID", Value: employeeID}}).Decode(&emp)

	if err != nil {
		return nil, err
	}

	return &emp, nil
}

func (r *EmployeeRepo) GetAllEmployees() ([]model.Employee, error) {
	result, err := r.MongoCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	var emps []model.Employee
	err = result.All(context.Background(), &emps)

	if err != nil {
		return nil, fmt.Errorf("results decode error %s", err.Error())
	}

	return emps, nil
}

func (r *EmployeeRepo) UpdateEmployee(employeeID string, newEmployee *model.Employee) (int, error) {
	result, err := r.MongoCollection.UpdateOne(context.Background(), 
					bson.D{{Key: "employeeID" , Value: employeeID}},
					bson.D{{Key: "$set", Value: newEmployee}})	
	
	if err != nil {
		return 0, err
	}

	return int(result.ModifiedCount), nil
}

func (r *EmployeeRepo) DeleteEmployee(employeeID string) (int, error) {
	result, err := r.MongoCollection.DeleteOne(context.Background(), 
					bson.D{{Key: "employeeID" , Value: employeeID}})	
	
	if err != nil {
		return 0, err
	}

	return int(result.DeletedCount), nil
}