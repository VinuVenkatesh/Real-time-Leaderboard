package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VinuVenkatesh/Real-time-Leaderboard/backend/services"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	// Handler logic will go here
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := services.RegisterUser(req.Username, req.Password); err != nil {
		http.Error(w, fmt.Sprintf("Registration failed: %v", err), http.StatusInternalServerError)
		return
	}

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Handler logic will go here
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	authenticated, err := services.AuthenticateUser(req.Username, req.Password)
	if err != nil {
		http.Error(w, fmt.Sprintf("Authentication failed: %v", err), http.StatusInternalServerError)
		return
	}
	if !authenticated {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}
}
