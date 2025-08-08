package repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/mohamedkaram400/go-crud-ops/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *EmployeeRepo) GetAllEmployees(skip int, limit int) ([]models.Employee, int, error) {

	totalCount, err := r.MongoCollection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		return nil, 0, err
	}

	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(skip))
	findOptions.SetSort(bson.M{"_id": -1})
 
	cursor, err := r.MongoCollection.Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		return nil, 0, err
	}

	defer cursor.Close(context.Background())


	var employees []models.Employee

	if 	err = cursor.All(context.Background(), &employees); err != nil {
		return nil, 0, fmt.Errorf("results decode error %s", err.Error())
	}

	return employees, int(totalCount), nil
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
