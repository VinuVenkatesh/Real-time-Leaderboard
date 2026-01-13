package services

import (
	"github.com/VinuVenkatesh/Real-time-Leaderboard/backend/db"
)

func GetAllSportsActivities() ([]string, error) {
	return db.GetActivities()
}
