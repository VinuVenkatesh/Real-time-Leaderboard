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
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := services.RegisterUser(request.Username, request.Password); err != nil {
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
	token, err := services.AuthenticateUser(req.Username, req.Password)
	if err != nil {
		http.Error(w, fmt.Sprintf("Authentication failed: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Token string `json:"token"`
	}{Token: token})
}
