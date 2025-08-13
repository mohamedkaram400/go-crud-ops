package main

import (
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
	"github.com/mohamedkaram400/go-crud-ops/repository"
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

	// 4. Get Employee Collection for employees
	employeeCollection := client.Database(config.GetDBName()).Collection(config.GetCollectionName())

	// 5. Create Repository layer
	employeeRepo := &repository.EmployeeRepo{MongoCollection: employeeCollection}
	
	// 6. Create service layer
	employeeService := &usecases.EmployeeService{Repo: employeeRepo}

	// 7. Create handler layer
	employeeHandler := &handlers.EmployeeHandler{Service: employeeService}
	
	// 8. Create router and register API routes
	router := mux.NewRouter()

	// 9. Add rate limiter validation
	router.Use(middlewares.RateLimiter(config.GetRateNumber(), time.Second)) 

	// 10. Add resource for auth module
	authRepo := &repository.AuthRepo{MongoCollection: employeeCollection}

	authService := &usecases.AuthService{Repo: authRepo}

	authHandler := &handlers.AuthHandler{Service: authService}

	// 11. Regsiter routes and give him api version 1
	api := router.PathPrefix("/api/v1").Subrouter()

	routes.RegisterAuthRoutes(api, authHandler)

	// Secure employee routes
	protected := api.NewRoute().Subrouter()
	
	protected.Use(middlewares.JWTAuth)

	routes.RegisterEmployeeRoutes(protected, employeeHandler)

	// 12. Start HTTP server
	StartServer(router)
}

func StartServer(router *mux.Router) {
	log.Println("🚀 Server is running on http://localhost:4444")
	log.Fatal(http.ListenAndServe(":4444", router))
}