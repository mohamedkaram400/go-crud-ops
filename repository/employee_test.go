package repository

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/mohamedkaram400/go-crud-ops/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoClient() *mongo.Client {
	err := godotenv.Load("../.env") 
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	mongoTestClient, err := mongo.Connect(context.Background(),
							options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatal("Error while connecting to mongo", err)
	}
	
	log.Println("Mongo successfully connected")

	err = mongoTestClient.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("ping failed", err)
	}

	println("ping success")

	return mongoTestClient
}

func TestMongoOperations(t *testing.T) {
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	emp1 := uuid.New().String()
	// emp2 = uuid.New().String()

	coll := mongoTestClient.Database("compnay").Collection("employeees")

	empRepo := EmployeeRepo{MongoCollection: coll}

	t.Run("Insert Employee 1", func (t *testing.T) {
		emp := models.Employee {
			Name: "Mohamed Karam",
			Department: "Backend Development",
			ID: emp1,
		}

		result, err := empRepo.InsertEmployee(&emp)

		if err != nil {
			t.Fatal("Insert 1 operation is failed")
		}

		t.Log("Insert 1 operation success", result)
	})
}