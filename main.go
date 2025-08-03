package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mohamedkaram400/go-crud-ops/config"
	"github.com/mohamedkaram400/go-crud-ops/db"
	"github.com/mohamedkaram400/go-crud-ops/handlers"
	"github.com/mohamedkaram400/go-crud-ops/internal/redis"
	"github.com/mohamedkaram400/go-crud-ops/middlewares"
	"github.com/mohamedkaram400/go-crud-ops/routes"
	"github.com/mohamedkaram400/go-crud-ops/usecases"
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

	// 3. Connect to Redis
	if err := redisclient.Init(); err != nil {
		log.Fatalf("❌ Failed to connect Redis: %v", err)
	}

	// 4. Get collection for employees
	collection := client.Database(config.GetDBName()).Collection(config.GetCollectionName())

	// 5. Create service layer
	employeeService := usecases.EmployeeService{MongoCollection: collection}

	// 6. Create handler layer
	empHandler := &handlers.EmployeeHandler{Service: &employeeService}
	
	// 7. Create router and register API routes
	router := mux.NewRouter()

	fmt.Println("ttt:", config.GetRateNumber())
	// 8. Add rate limiter validation
	router.Use(middlewares.RateLimiter(config.GetRateNumber(), 10 * time.Second)) 

	routes.RegisterAPIV1Routes(router, empHandler)

	// 9. Start HTTP server
	StartServer(router)
}

func StartServer(router *mux.Router) {
	log.Println("🚀 Server is running on http://localhost:4444")
	log.Fatal(http.ListenAndServe(":4444", router))
}