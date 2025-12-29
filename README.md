# Go Employee Service – REST & gRPC

A scalable **Employee Management Service** built with **Go**, supporting both **REST API and gRPC** communication.

This project focuses on **clean architecture, security, scalability, and real-world backend concerns** such as authentication, rate limiting, pagination, and protocol flexibility.

---

## 🧱 Tech Stack

- Go (Golang)
- REST API (net/http)
- gRPC
- MongoDB
- Redis
- JWT Authentication
- Clean Architecture (Repository & Service layers)

---

## 🎯 Project Goal

Build a **production-style CRUD service** that:
- Manages employees
- Supports REST and gRPC side-by-side
- Implements secure authentication
- Handles high traffic safely
- Scales cleanly with well-defined layers

---

## ✨ Features

### 1️⃣ Employee CRUD
- Create, read, update, delete employees
- Pagination support (`page`, `limit`)
- Scalable data access layer using MongoDB
- Repository interfaces for easy extension and testing

---

### 2️⃣ Authentication & Security
- JWT-based authentication
- Login & registration for employees
- Refresh token endpoint
- Redis-backed token management
- Secure token invalidation
- Real client IP extraction from requests

---

### 3️⃣ Rate Limiting
- Redis-based rate limiter
- Protects against DoS attacks
- Limits requests per IP
- Designed to work consistently across REST endpoints

---

### 4️⃣ REST & gRPC Support
- REST API for standard HTTP clients
- gRPC protocol for high-performance internal communication
- Shared business logic across both protocols
- Demonstrates protocol-agnostic service design

---

## 🧠 Architecture

- Clean separation of concerns:
  - Handlers (REST / gRPC)
  - Services
  - Repositories
  - Data layer
- Repository interfaces for scalability
- Business logic isolated from transport layer
- Redis used for:
  - Rate limiting
  - JWT refresh token handling

---

## 🧪 Reliability Considerations

- Proper pagination handling
- Safe token refresh flow
- Rate limiting to prevent abuse
- Consistent data handling across REST and gRPC
- Designed to scale with additional services

---

## 🚀 Getting Started

### Prerequisites
- Go 1.21+
- MongoDB
- Redis

### Run the project

```bash
go install github.com/air-verse/air@latest
export PATH=$PATH:$(go env GOPATH)/bin
go mod tidy
air
```

## 👤 Author
### Mohamed Karam
Backend Developer — Laravel & Go

---

## 🛡️ License

This project is licensed under the [MIT License](LICENSE). You are free to use, modify, and share this project with proper attribution.
