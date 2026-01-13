package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/VinuVenkatesh/Real-time-Leaderboard/backend/services"
)

func GetActivitiesHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation here
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	activities, err := services.GetAllSportsActivities()
	if err != nil {
		http.Error(w, "Failed to retrieve activities", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Activities []string `json:"activities"`
	}{Activities: activities})

}
