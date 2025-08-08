package usecases

import (
	"github.com/mohamedkaram400/go-crud-ops/models"
	"github.com/mohamedkaram400/go-crud-ops/helpers"
	"github.com/mohamedkaram400/go-crud-ops/requests"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"errors"
	"context"
)

type AuthService struct {
	MongoCollection *mongo.Collection
}

// func (svc *AuthService) register() (*models.Employee, error) {

// }

func (svc *AuthService) Login(req *requests.LoginRequest) (*models.Employee, error) {
	var emp models.Employee
	filter := bson.M{"username": req.UserName}
	
	err := svc.MongoCollection.FindOne(context.TODO(), filter).Decode(&emp)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	if err := helpers.CheckPassword(emp.Password, req.Password); err != nil {
		return nil, errors.New("invalid username or password")
	}

	return &emp, nil
}

func (svc *AuthService) Logout(employeeID string) error {
	// In JWT auth, logout is usually stateless
	// Optionally, store blacklisted tokens in Redis
	return nil
}