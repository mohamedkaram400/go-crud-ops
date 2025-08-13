package repository

import (

	"go.mongodb.org/mongo-driver/bson"
	"github.com/mohamedkaram400/go-crud-ops/models"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
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

// func (r *AuthRepo) Register(data) (*models.Employee, string, error) {

// 	return employee, "User register success", nil
// }

// func (r *AuthRepo) Login(userName string, password string) (*models.Employee, string, error) {

// 	return employee, "User login success", nil
// }

// func (r *AuthRepo) Logout() (string, error) {

// 	return "User logout success", nil
// }

