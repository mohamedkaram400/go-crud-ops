package main

import (
	"log"
	"net/http"
	"time"
	"net"

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
    "google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
	pb "github.com/mohamedkaram400/go-crud-ops/proto"
)

func main() {

	// 1. Load env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// 2. Connect DBs
	client, err := db.ConnectMongo(config.GetMongoURI())	// MongoDB
	if err != nil {
		log.Fatal("❌ Failed to connect Mongo:", err)
	}
	if err := redisclient.Init(); err != nil {				// Redis
		log.Fatalf("❌ Failed to connect Redis: %v", err)
	}
	
	// 3. Init repositories
	employeeCollection := client.Database(config.GetDBName()).Collection(config.GetCollectionName())
	employeeRepo := &repository.EmployeeRepo{MongoCollection: employeeCollection}
	authRepo := &repository.AuthRepo{MongoCollection: employeeCollection}

	// 4. Init services
	employeeService := &usecases.EmployeeService{Repo: employeeRepo}
	authService := &usecases.AuthService{Repo: authRepo}

	// 5. Init handlers
	authHandler := &handlers.AuthHandler{Service: authService}
	authGRPCHandler := &handlers.AuthGRPCHandler{Service: authService}
	
	employeeHandler := &handlers.EmployeeHandler{Service: employeeService}
	employeeGRPCHandler := &handlers.EmployeeGRPCHandler{Service: employeeService}
		
	// 6. Router
	router := mux.NewRouter()
	router.Use(middlewares.RateLimiter(config.GetRateNumber(), time.Second))

	// 7. API v1
	api := router.PathPrefix("/api/v1").Subrouter()
	routes.RegisterAuthRoutes(api, authHandler)
	routes.RegisterEmployeeRoutes(api, employeeHandler)

	// 8. Run REST + gRPC in goroutines
	go startREST(router)
	go startGRPC(authGRPCHandler, employeeGRPCHandler)

	// Block forever
	select {}
}

func startREST(router *mux.Router) {
	log.Println("🌍 REST server running at http://localhost:4444")
	log.Fatal(http.ListenAndServe(":4444", router))
}

func startGRPC(authGRPCHandler *handlers.AuthGRPCHandler, employeeGRPCHandler *handlers.EmployeeGRPCHandler) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, authGRPCHandler)
	pb.RegisterEmployeeServiceServer(grpcServer, employeeGRPCHandler)

    reflection.Register(grpcServer)

	log.Println("⚡ gRPC server running at :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ failed to serve: %v", err)
	}
}