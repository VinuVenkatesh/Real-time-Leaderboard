package routes

import (
	"net/http"

	"github.com/VinuVenkatesh/Real-time-Leaderboard/backend/handlers"
)

func ActivityRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/listactivities", handlers.GetActivitiesHandler)
}
