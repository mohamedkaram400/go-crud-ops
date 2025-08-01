package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohamedkaram400/go-crud-ops/usecases"
	"github.com/mohamedkaram400/go-crud-ops/handlers"
	"github.com/mohamedkaram400/go-crud-ops/routes"
	"github.com/mohamedkaram400/go-crud-ops/db"
	"github.com/mohamedkaram400/go-crud-ops/config"
	"github.com/joho/godotenv"
)


func main() {

	// 1. Load env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// 2. Connect to MongoDB
	client, err := db.ConnectMongo(config.GetMongoURI())
	if err != nil {
		log.Fatal("❌ Failed to connect Mongo:", err)
	}

	// 3. Get collection for employees
	collection := client.Database(config.GetDBName()).Collection(config.GetCollectionName())

	// 4. Create service layer
	employeeService := usecases.EmployeeService{MongoCollection: collection}

	// 5. Create handler layer
	empHandler := &handlers.EmployeeHandler{Service: &employeeService}

	// 6. Create router and register API routes
	router := mux.NewRouter()
	routes.RegisterAPIV1Routes(router, empHandler)

	// 7. Start HTTP server
	StartServer(router)
}

func StartServer(router *mux.Router) {
	log.Println("🚀 Server is running on http://localhost:4444")
	log.Fatal(http.ListenAndServe(":4444", router))
}