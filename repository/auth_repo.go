package repository

import (
	"context"

	"github.com/mohamedkaram400/go-crud-ops/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepo struct {
	MongoCollection *mongo.Collection
}


func (r *AuthRepo) GetEmployeeByUsername(ctx context.Context, username string) (*models.Employee, error) {
	var emp models.Employee
	err := r.MongoCollection.FindOne(ctx, bson.M{"username": username}).Decode(&emp)
	if err != nil {
		return nil, err
	}
	return &emp, nil
}

func (r *AuthRepo) Register(employee *models.Employee) (*models.Employee, string, error) {
	_, err := r.MongoCollection.InsertOne(context.Background(), employee)
	if err != nil {
		return nil, "Error happened when creating new employee", err
	}
	return employee, "User register success", nil
}

func (r *AuthRepo) Logout(employeeID string) (string, error) {

	return "User logout success", nil
}

