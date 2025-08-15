package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mohamedkaram400/go-crud-ops/middlewares"
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

func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req requests.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	emp, msg, err := h.Service.Register(r.Context(), &req)
	if err != nil {
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  msg,
		"employee": emp,
	})
}

func (h *AuthHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Get employee ID from context (set in JWT middleware)
	empID, ok := r.Context().Value(middlewares.EmployeeIDKey).(string)
	if !ok || empID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	msg, err := h.Service.Logout(empID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": msg})
}