package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mohamedkaram400/go-crud-ops/requests"
	"github.com/mohamedkaram400/go-crud-ops/usecases"
)

type AuthHandler struct {
	Service *usecases.AuthService
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req requests.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	token, err := h.Service.Login(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (svc *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &PaginatedResult{}
	defer json.NewEncoder(w).Encode(res)
}

func (h *AuthHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Logged out"}`))
}